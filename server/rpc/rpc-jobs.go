package rpc

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"
	"errors"

	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/commonpb"
	"github.com/bishopfox/sliver/server/c2"
	"github.com/bishopfox/sliver/server/configs"
	"github.com/bishopfox/sliver/server/core"
	"github.com/bishopfox/sliver/server/db"
	"github.com/bishopfox/sliver/server/db/models"
)

const (
	defaultMTLSPort    = 4444
	defaultWGPort      = 53
	defaultWGNPort     = 8888
	defaultWGKeyExPort = 1337
	defaultDNSPort     = 53
	defaultHTTPPort    = 80
	defaultHTTPSPort   = 443
)

var (
	// ErrInvalidPort - Invalid TCP port number
	ErrInvalidPort = errors.New("invalid listener port")
)

// GetJobs - List jobs
func (rpc *Server) GetJobs(ctx context.Context, _ *commonpb.Empty) (*clientpb.Jobs, error) {
	jobs := &clientpb.Jobs{
		Active: []*clientpb.Job{},
	}
	for _, job := range core.Jobs.All() {
		jobs.Active = append(jobs.Active, &clientpb.Job{
			ID:          uint32(job.ID),
			Name:        job.Name,
			Description: job.Description,
			Protocol:    job.Protocol,
			Port:        uint32(job.Port),
			Domains:     job.Domains,
		})
	}
	return jobs, nil
}

// KillJob - Kill a server-side job
func (rpc *Server) KillJob(ctx context.Context, kill *clientpb.KillJobReq) (*clientpb.KillJob, error) {
	job := core.Jobs.Get(int(kill.ID))
	killJob := &clientpb.KillJob{}
	var err error = nil
	if job != nil {
		job.JobCtrl <- true
		killJob.ID = uint32(job.ID)
		killJob.Success = true
		if job.PersistentID != "" {
			configs.GetServerConfig().RemoveJob(job.PersistentID)
		}
	} else {
		killJob.Success = false
		err = errors.New("invalid Job ID")
	}
	return killJob, err
}

// StartMTLSListener - Start an MTLS listener
func (rpc *Server) StartMTLSListener(ctx context.Context, req *clientpb.MTLSListenerReq) (*clientpb.ListenerJob, error) {

	if 65535 <= req.Port {
		return nil, ErrInvalidPort
	}
	listenPort := uint16(defaultMTLSPort)
	if req.Port != 0 {
		listenPort = uint16(req.Port)
	}

	job, err := c2.StartMTLSListenerJob(req.Host, listenPort)
	if err != nil {
		return nil, err
	}

	listenerJob := &clientpb.ListenerJob{
		JobID:    uint32(job.ID),
		Type:     "mtls",
		MTLSConf: req,
	}
	listenerModel := models.ListenerJobFromProtobuf(listenerJob)
	db.HTTPC2ListenerSave(listenerModel)

	return &clientpb.ListenerJob{JobID: uint32(job.ID)}, nil
}

// StartWGListener - Start a Wireguard listener
func (rpc *Server) StartWGListener(ctx context.Context, req *clientpb.WGListenerReq) (*clientpb.ListenerJob, error) {

	if 65535 <= req.Port || 65535 <= req.NPort || 65535 <= req.KeyPort {
		return nil, ErrInvalidPort
	}
	listenPort := uint16(defaultWGPort)
	if req.Port != 0 {
		listenPort = uint16(req.Port)
	}

	nListenPort := uint16(defaultWGNPort)
	if req.NPort != 0 {
		nListenPort = uint16(req.NPort)
	}

	keyExchangeListenPort := uint16(defaultWGKeyExPort)
	if req.NPort != 0 {
		keyExchangeListenPort = uint16(req.KeyPort)
	}

	job, err := c2.StartWGListenerJob(listenPort, nListenPort, keyExchangeListenPort)
	if err != nil {
		return nil, err
	}

	listenerJob := &clientpb.ListenerJob{
		JobID:  uint32(job.ID),
		Type:   "wg",
		WGConf: req,
	}
	listenerModel := models.ListenerJobFromProtobuf(listenerJob)
	db.HTTPC2ListenerSave(listenerModel)

	return &clientpb.ListenerJob{JobID: uint32(job.ID)}, nil
}

// StartDNSListener - Start a DNS listener TODO: respect request's Host specification
func (rpc *Server) StartDNSListener(ctx context.Context, req *clientpb.DNSListenerReq) (*clientpb.ListenerJob, error) {
	if 65535 <= req.Port {
		return nil, ErrInvalidPort
	}
	listenPort := uint16(defaultDNSPort)
	if req.Port != 0 {
		listenPort = uint16(req.Port)
	}

	job, err := c2.StartDNSListenerJob(req.Host, listenPort, req.Domains, req.Canaries, req.EnforceOTP)
	if err != nil {
		return nil, err
	}

	// TODO save listener to db

	return &clientpb.ListenerJob{JobID: uint32(job.ID)}, nil
}

// StartHTTPSListener - Start an HTTPS listener
func (rpc *Server) StartHTTPSListener(ctx context.Context, req *clientpb.HTTPListenerReq) (*clientpb.ListenerJob, error) {
	if 65535 <= req.Port {
		return nil, ErrInvalidPort
	}
	if req.Port == 0 {
		req.Port = defaultHTTPSPort
	}

	job, err := c2.StartHTTPListenerJob(req)
	if err != nil {
		return nil, err
	}

	listenerJob := &clientpb.ListenerJob{
		JobID:    uint32(job.ID),
		Type:     "http",
		HTTPConf: req,
	}
	listenerModel := models.ListenerJobFromProtobuf(listenerJob)
	db.HTTPC2ListenerSave(listenerModel)

	return &clientpb.ListenerJob{JobID: uint32(job.ID)}, nil
}

// StartHTTPListener - Start an HTTP listener
func (rpc *Server) StartHTTPListener(ctx context.Context, req *clientpb.HTTPListenerReq) (*clientpb.ListenerJob, error) {
	if 65535 <= req.Port {
		return nil, ErrInvalidPort
	}
	if req.Port == 0 {
		req.Port = defaultHTTPPort
	}

	job, err := c2.StartHTTPListenerJob(req)
	if err != nil {
		return nil, err
	}

	listenerJob := &clientpb.ListenerJob{
		JobID:    uint32(job.ID),
		Type:     "http",
		HTTPConf: req,
	}
	listenerModel := models.ListenerJobFromProtobuf(listenerJob)
	db.HTTPC2ListenerSave(listenerModel)

	return &clientpb.ListenerJob{JobID: uint32(job.ID)}, nil
}
