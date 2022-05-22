// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the MIT License. See License-MIT.txt in the project root for license information.

package server

import (
	"encoding/json"
	"testing"

	"github.com/gitpod-io/gitpod/installer/pkg/common"
	"github.com/gitpod-io/gitpod/installer/pkg/config/v1"
	"github.com/gitpod-io/gitpod/installer/pkg/config/v1/experimental"
	"github.com/gitpod-io/gitpod/installer/pkg/config/versions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func TestConfigMap(t *testing.T) {
	type Expectation struct {
		EnableLocalApp                    bool
		RunDbDeleter                      bool
		DisableDynamicAuthProviderLogin   bool
		DisableWorkspaceGarbageCollection bool
		DefaultBaseImageRegistryWhiteList []string
		WorkspaceImage                    string
		JWTSecret                         string
		SessionSecret                     string
		BlockedRepositories               []experimental.BlockedRepository
		GitHubApp                         experimental.GithubApp
	}

	expectation := Expectation{
		EnableLocalApp:                    true,
		DisableDynamicAuthProviderLogin:   true,
		RunDbDeleter:                      false,
		DisableWorkspaceGarbageCollection: true,
		DefaultBaseImageRegistryWhiteList: []string{"some-registry"},
		WorkspaceImage:                    "some-workspace-image",
		JWTSecret:                         "some-jwt-secret",
		SessionSecret:                     "some-session-secret",
		BlockedRepositories: []experimental.BlockedRepository{{
			UrlRegExp: "https://github.com/some-user/some-bad-repo",
			BlockUser: true,
		}},
		GitHubApp: experimental.GithubApp{
			AppId:           123,
			AuthProviderId:  "some-auth-provider-id",
			BaseUrl:         "some-base-url",
			CertPath:        "some-cert-path",
			Enabled:         true,
			LogLevel:        "some-log-level",
			MarketplaceName: "some-marketplace-name",
			WebhookSecret:   "some-webhook-secret",
			CertSecretName:  "some-cert-secret-name",
		},
	}

	ctx, err := common.NewRenderContext(config.Config{
		Experimental: &experimental.Config{
			WebApp: &experimental.WebAppConfig{
				Server: &experimental.ServerConfig{
					DisableDynamicAuthProviderLogin:   expectation.DisableDynamicAuthProviderLogin,
					EnableLocalApp:                    pointer.Bool(expectation.EnableLocalApp),
					RunDbDeleter:                      pointer.Bool(expectation.RunDbDeleter),
					DisableWorkspaceGarbageCollection: expectation.DisableWorkspaceGarbageCollection,
					DefaultBaseImageRegistryWhiteList: expectation.DefaultBaseImageRegistryWhiteList,
					WorkspaceDefaults: experimental.WorkspaceDefaults{
						WorkspaceImage: expectation.WorkspaceImage,
					},
					OAuthServer: experimental.OAuthServer{
						JWTSecret: expectation.JWTSecret,
					},
					Session: experimental.Session{
						Secret: expectation.SessionSecret,
					},
					GithubApp:           &expectation.GitHubApp,
					BlockedRepositories: expectation.BlockedRepositories,
				},
			},
		},
	}, versions.Manifest{}, "test_namespace")

	require.NoError(t, err)
	objs, err := configmap(ctx)
	if err != nil {
		t.Errorf("failed to generate configmap: %s\n", err)
	}

	configmap, ok := objs[0].(*corev1.ConfigMap)
	if !ok {
		t.Fatalf("rendering configmap did not return a configMap")
		return
	}

	configJson, ok := configmap.Data["config.json"]
	if ok == false {
		t.Errorf("no %q key found in configmap data", "config.json")
	}

	var config ConfigSerialized
	if err := json.Unmarshal([]byte(configJson), &config); err != nil {
		t.Errorf("failed to unmarshal config json: %s", err)
	}

	actual := Expectation{
		DisableDynamicAuthProviderLogin:   config.DisableDynamicAuthProviderLogin,
		EnableLocalApp:                    config.EnableLocalApp,
		RunDbDeleter:                      config.RunDbDeleter,
		DisableWorkspaceGarbageCollection: config.WorkspaceGarbageCollection.Disabled,
		DefaultBaseImageRegistryWhiteList: config.DefaultBaseImageRegistryWhitelist,
		WorkspaceImage:                    config.WorkspaceDefaults.WorkspaceImage,
		JWTSecret:                         config.OAuthServer.JWTSecret,
		SessionSecret:                     config.Session.Secret,
		BlockedRepositories: func(config ConfigSerialized) []experimental.BlockedRepository {
			var blockedRepos []experimental.BlockedRepository
			for _, repo := range config.BlockedRepositories {
				blockedRepos = append(blockedRepos, experimental.BlockedRepository{
					UrlRegExp: repo.UrlRegExp,
					BlockUser: repo.BlockUser,
				})
			}
			return blockedRepos
		}(config),
		GitHubApp: experimental.GithubApp{
			AppId:           config.GitHubApp.AppId,
			AuthProviderId:  config.GitHubApp.AuthProviderId,
			BaseUrl:         config.GitHubApp.BaseUrl,
			CertPath:        config.GitHubApp.CertPath,
			Enabled:         config.GitHubApp.Enabled,
			LogLevel:        config.GitHubApp.LogLevel,
			MarketplaceName: config.GitHubApp.MarketplaceName,
			WebhookSecret:   config.GitHubApp.WebhookSecret,
			CertSecretName:  config.GitHubApp.CertSecretName,
		},
	}

	assert.Equal(t, expectation, actual)
}

func TestInvalidBlockedRepositoryRegularExpressions(t *testing.T) {
	const invalidRegexp = "["

	ctx, err := common.NewRenderContext(config.Config{
		Experimental: &experimental.Config{
			WebApp: &experimental.WebAppConfig{
				Server: &experimental.ServerConfig{
					BlockedRepositories: []experimental.BlockedRepository{{
						UrlRegExp: invalidRegexp,
						BlockUser: false,
					}},
				},
			},
		},
	}, versions.Manifest{}, "test_namespace")
	require.NoError(t, err)

	_, err = configmap(ctx)

	require.Error(t, err, "expected to fail when rendering configmap with invalid blocked repo regexp %q", invalidRegexp)
}
