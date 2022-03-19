/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { User } from "@gitpod/gitpod-protocol";
import { skipIfEnvVarNotSet } from "@gitpod/gitpod-protocol/lib/util/skip-if";
import { Container, ContainerModule } from "inversify";
import { suite, test, timeout } from "mocha-typescript";
import { expect } from "chai";
import { GitpodHostUrl } from "@gitpod/gitpod-protocol/lib/util/gitpod-host-url";
import { BitbucketServerFileProvider } from "./bitbucket-server-file-provider";
import { AuthProviderParams } from "../auth/auth-provider";
import { BitbucketServerContextParser } from "./bitbucket-server-context-parser";
import { BitbucketServerTokenHelper } from "./bitbucket-server-token-handler";
import { TokenService } from "../user/token-service";
import { Config } from "../config";
import { TokenProvider } from "../user/token-provider";
import { BitbucketServerApi } from "./bitbucket-server-api";
import { HostContextProvider } from "../auth/host-context-provider";

@suite(timeout(10000), skipIfEnvVarNotSet("GITPOD_TEST_TOKEN_BITBUCKET_SERVER"))
class TestBitbucketServerContextParser {
    protected parser: BitbucketServerContextParser;
    protected user: User;

    static readonly AUTH_HOST_CONFIG: Partial<AuthProviderParams> = {
        id: "MyBitbucketServer",
        type: "BitbucketServer",
        host: "bitbucket.gitpod-self-hosted.com",
    };

    public before() {
        const container = new Container();
        container.load(
            new ContainerModule((bind, unbind, isBound, rebind) => {
                bind(BitbucketServerFileProvider).toSelf().inSingletonScope();
                bind(BitbucketServerContextParser).toSelf().inSingletonScope();
                bind(AuthProviderParams).toConstantValue(TestBitbucketServerContextParser.AUTH_HOST_CONFIG);
                bind(BitbucketServerTokenHelper).toSelf().inSingletonScope();
                bind(TokenService).toConstantValue({
                    createGitpodToken: async () => ({ token: { value: "foobar123-token" } }),
                } as any);
                bind(Config).toConstantValue({
                    hostUrl: new GitpodHostUrl(),
                });
                bind(TokenProvider).toConstantValue(<TokenProvider>{
                    getTokenForHost: async () => {
                        return {
                            value: process.env["GITPOD_TEST_TOKEN_BITBUCKET_SERVER"] || "undefined",
                            scopes: [],
                        };
                    },
                    getFreshPortAuthenticationToken: undefined as any,
                });
                bind(BitbucketServerApi).toSelf().inSingletonScope();
                bind(HostContextProvider).toConstantValue({
                    get: (hostname: string) => {
                        authProvider: {
                            ("BBS");
                        }
                    },
                });
            }),
        );
        this.parser = container.get(BitbucketServerContextParser);
        this.user = {
            creationDate: "",
            id: "user1",
            identities: [
                {
                    authId: "user1",
                    authName: "AlexTugarev",
                    authProviderId: "MyBitbucketServer",
                },
            ],
        };
    }

    @test async test_tree_context_01() {
        const result = await this.parser.handle(
            {},
            this.user,
            "https://bitbucket.gitpod-self-hosted.com/projects/FOO/repos/repo123",
        );
        console.log(JSON.stringify(result, undefined, "  "));
        expect(result).to.deep.include({
            ref: "master",
            refType: "branch",
            revision: "d00195289e73f0257a579dd68bf76e22480fd4ba",
            path: "",
            isFile: false,
            repository: {
                host: "bitbucket.gitpod-self-hosted.com",
                owner: "FOO",
                name: "repo123",
                cloneUrl: "https://bitbucket.gitpod-self-hosted.com/scm/foo/repo123.git",
                webUrl: "https://bitbucket.gitpod-self-hosted.com/projects/FOO/repos/repo123",
                defaultBranch: "master",
                private: true,
                repoKind: "projects",
            },
            title: "FOO/repo123 - master",
        });
    }

    @test async test_tree_context_02() {
        const result = await this.parser.handle(
            {},
            this.user,
            "https://bitbucket.gitpod-self-hosted.com/scm/foo/repo123.git",
        );
        console.log(JSON.stringify(result, undefined, "  "));
        expect(result).to.deep.include({
            ref: "master",
            refType: "branch",
            revision: "d00195289e73f0257a579dd68bf76e22480fd4ba",
            path: "",
            isFile: false,
            repository: {
                host: "bitbucket.gitpod-self-hosted.com",
                owner: "FOO",
                name: "repo123",
                cloneUrl: "https://bitbucket.gitpod-self-hosted.com/scm/foo/repo123.git",
                webUrl: "https://bitbucket.gitpod-self-hosted.com/projects/FOO/repos/repo123",
                defaultBranch: "master",
                private: true,
                repoKind: "projects",
            },
            title: "foo/repo123 - master",
        });
    }
}

module.exports = new TestBitbucketServerContextParser();
