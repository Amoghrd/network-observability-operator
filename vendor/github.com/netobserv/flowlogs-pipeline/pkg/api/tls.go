/*
 * Copyright (C) 2022 IBM, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package api

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

type ClientTLS struct {
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify,omitempty" json:"insecureSkipVerify,omitempty" doc:"skip client verifying the server's certificate chain and host name"`
	CACertPath         string `yaml:"caCertPath,omitempty" json:"caCertPath,omitempty" doc:"path to the CA certificate"`
	UserCertPath       string `yaml:"userCertPath,omitempty" json:"userCertPath,omitempty" doc:"path to the user certificate"`
	UserKeyPath        string `yaml:"userKeyPath,omitempty" json:"userKeyPath,omitempty" doc:"path to the user private key"`
}

type ServerTLS struct {
	CertPath       string `yaml:"certPath,omitempty" json:"certPath,omitempty" doc:"path to the certificate to enable TLS"`
	PrivateKeyPath string `yaml:"privateKeyPath,omitempty" json:"privateKeyPath,omitempty" doc:"path to the private key to enable TLS"`
}

func (c *ClientTLS) Build() (*tls.Config, error) {
	if c.InsecureSkipVerify {
		return &tls.Config{
			InsecureSkipVerify: true,
		}, nil
	}
	if c.CACertPath != "" {
		caCert, err := ioutil.ReadFile(c.CACertPath)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caCert)
		tlsConfig := &tls.Config{RootCAs: pool}

		if c.UserCertPath != "" && c.UserKeyPath != "" {
			userCert, err := ioutil.ReadFile(c.UserCertPath)
			if err != nil {
				return nil, err
			}
			userKey, err := ioutil.ReadFile(c.UserKeyPath)
			if err != nil {
				return nil, err
			}
			pair, err := tls.X509KeyPair([]byte(userCert), []byte(userKey))
			if err != nil {
				return nil, err
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
			return tlsConfig, nil
		}
	}
	return nil, nil
}

func (c *ServerTLS) Build() (*tls.Config, string, string) {
	// Clients must use TLS 1.2 or higher
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	return tlsConfig, c.PrivateKeyPath, c.CertPath
}
