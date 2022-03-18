/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { IDEOption, IDEOptions } from "@gitpod/gitpod-protocol/lib/ide-protocol";
import { useContext, useEffect, useState } from "react";
import InfoBox from "../components/InfoBox";
import { PageWithSubMenu } from "../components/PageWithSubMenu";
import PillLabel from "../components/PillLabel";
import SelectableCard from "../components/SelectableCard";
import SelectableCardSolid from "../components/SelectableCardSolid";
import Tooltip from "../components/Tooltip";
import { getGitpodService } from "../service/service";
import { ThemeContext } from "../theme-context";
import { UserContext } from "../user-context";
import settingsMenu from "./settings-menu";
import CheckBox from "../components/CheckBox";
import { trackEvent } from "../Analytics";

type Theme = "light" | "dark" | "system";

const DefaultBrowserIDE = "code";

export default function Preferences() {
    const { user } = useContext(UserContext);
    const { setIsDark } = useContext(ThemeContext);

    const updateUserIDEInfo = async (selectedIde: string, useLatestVersion: boolean) => {
        // Always start vscode browser ide
        const defaultIde = DefaultBrowserIDE;
        const useDesktopIde = selectedIde !== DefaultBrowserIDE;
        const desktopIde = useDesktopIde ? selectedIde : undefined;
        const additionalData = user?.additionalData ?? {};
        const settings = additionalData.ideSettings ?? {};
        settings.useDesktopIde = useDesktopIde;
        settings.defaultIde = defaultIde;
        settings.defaultDesktopIde = desktopIde;
        settings.useLatestVersion = useLatestVersion;
        additionalData.ideSettings = settings;
        getGitpodService()
            .server.trackEvent({
                event: "ide_configuration_changed",
                properties: {
                    useDesktopIde,
                    defaultIde,
                    defaultDesktopIde: desktopIde,
                    useLatestVersion,
                },
            })
            .then()
            .catch(console.error);
        await getGitpodService().server.updateLoggedInUser({ additionalData });
    };

    const [defaultIde, setDefaultIde] = useState<string>(user?.additionalData?.ideSettings?.defaultIde || "");

    const [defaultDesktopIde, setDefaultDesktopIde] = useState<string>(
        user?.additionalData?.ideSettings?.defaultDesktopIde || "",
    );
    const actuallySetDefaultDesktopIde = async (value: string) => {
        await updateUserIDEInfo(value, useLatestVersion);
        setDefaultDesktopIde(value);
    };

    const [useLatestVersion, setUseLatestVersion] = useState<boolean>(
        user?.additionalData?.ideSettings?.useLatestVersion ?? false,
    );
    const actuallySetUseLatestVersion = async (value: boolean) => {
        await updateUserIDEInfo(defaultDesktopIde, value);
        setUseLatestVersion(value);
    };

    const [ideOptions, setIdeOptions] = useState<IDEOptions | undefined>(undefined);
    useEffect(() => {
        (async () => {
            const ideopts = await getGitpodService().server.getIDEOptions();
            setIdeOptions(ideopts);
            if (!defaultIde || ideopts.options[defaultIde] == null) {
                setDefaultIde(ideopts.defaultIde);
            }
            if (!defaultDesktopIde || ideopts.options[defaultDesktopIde] == null) {
                setDefaultDesktopIde(ideopts.defaultDesktopIde);
            }
        })();
    }, []);

    const [theme, setTheme] = useState<Theme>(localStorage.theme || "system");
    const actuallySetTheme = (theme: Theme) => {
        if (theme === "dark" || theme === "light") {
            localStorage.theme = theme;
        } else {
            localStorage.removeItem("theme");
        }
        const isDark =
            localStorage.theme === "dark" ||
            (localStorage.theme !== "light" && window.matchMedia("(prefers-color-scheme: dark)").matches);
        setIsDark(isDark);
        setTheme(theme);
    };

    const allIdeOptions = ideOptions && orderedIdeOptions(ideOptions);

    const [dotfileRepo, setDotfileRepo] = useState<string>(user?.additionalData?.dotfileRepo || "");
    const actuallySetDotfileRepo = async (value: string) => {
        const additionalData = user?.additionalData || {};
        const prevDotfileRepo = additionalData.dotfileRepo;
        additionalData.dotfileRepo = value;
        await getGitpodService().server.updateLoggedInUser({ additionalData });
        trackEvent("dotfile_repo_changed", {
            previous: prevDotfileRepo,
            current: value,
        });
    };

    return (
        <div>
            <PageWithSubMenu subMenu={settingsMenu} title="Preferences" subtitle="Configure user preferences.">
                {ideOptions && (
                    <>
                        {/* <div>{JSON.stringify(user?.additionalData?.ideSettings, null, 4)}</div> */}
                        {allIdeOptions && (
                            <>
                                <h3>Editor</h3>
                                <p className="text-base text-gray-500 dark:text-gray-400">
                                    Choose the editor for opening workspaces.
                                </p>
                                <div className="my-4 gap-4 flex flex-wrap max-w-2xl">
                                    {allIdeOptions.map(([id, option]) => {
                                        const selected = defaultDesktopIde === id;
                                        const onSelect = () => actuallySetDefaultDesktopIde(id);
                                        return renderIdeOption(option, selected, onSelect);
                                    })}
                                </div>
                                {ideOptions.options[defaultDesktopIde]?.notes && (
                                    <InfoBox className="my-5 max-w-2xl">
                                        <ul>
                                            {ideOptions.options[defaultDesktopIde].notes?.map((x, idx) => (
                                                <li className={idx > 0 ? "mt-2" : ""}>{x}</li>
                                            ))}
                                        </ul>
                                    </InfoBox>
                                )}
                                <p className="text-left w-full text-gray-500">
                                    The <strong>JetBrains desktop IDEs</strong> are currently in beta.{" "}
                                    <a
                                        href="https://github.com/gitpod-io/gitpod/issues/6576"
                                        target="gitpod-feedback-issue"
                                        rel="noopener"
                                        className="gp-link"
                                    >
                                        Send feedback
                                    </a>{" "}
                                    ·{" "}
                                    <a
                                        href="https://www.gitpod.io/docs/integrations/jetbrains"
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="gp-link"
                                    >
                                        Documentation
                                    </a>
                                </p>
                            </>
                        )}
                        <CheckBox
                            title="Latest Release"
                            desc="Use the latest version for each editor. Insiders for VS Code, EAP for JetBrains IDEs."
                            checked={useLatestVersion}
                            onChange={(e) => actuallySetUseLatestVersion(e.target.checked)}
                        />
                    </>
                )}
                <h3 className="mt-12">Theme</h3>
                <p className="text-base text-gray-500 dark:text-gray-400">Early bird or night owl? Choose your side.</p>
                <div className="mt-4 space-x-4 flex">
                    <SelectableCard
                        className="w-36 h-32"
                        title="Light"
                        selected={theme === "light"}
                        onClick={() => actuallySetTheme("light")}
                    >
                        <div className="flex-grow flex justify-center items-end">
                            <svg className="h-16" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 108 64">
                                <rect width="68" height="40" x="40" fill="#C4C4C4" rx="8" />
                                <rect width="32" height="16" fill="#C4C4C4" rx="8" />
                                <rect width="32" height="16" y="24" fill="#C4C4C4" rx="8" />
                                <rect width="32" height="16" y="48" fill="#C4C4C4" rx="8" />
                            </svg>
                        </div>
                    </SelectableCard>
                    <SelectableCard
                        className="w-36 h-32"
                        title="Dark"
                        selected={theme === "dark"}
                        onClick={() => actuallySetTheme("dark")}
                    >
                        <div className="flex-grow flex justify-center items-end">
                            <svg className="h-16" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 108 64">
                                <rect width="68" height="40" x="40" fill="#737373" rx="8" />
                                <rect width="32" height="16" fill="#737373" rx="8" />
                                <rect width="32" height="16" y="24" fill="#737373" rx="8" />
                                <rect width="32" height="16" y="48" fill="#737373" rx="8" />
                            </svg>
                        </div>
                    </SelectableCard>
                    <SelectableCard
                        className="w-36 h-32"
                        title="System"
                        selected={theme === "system"}
                        onClick={() => actuallySetTheme("system")}
                    >
                        <div className="flex-grow flex justify-center items-end">
                            <svg className="h-16" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 108 64">
                                <rect width="68" height="40" x="40" fill="#C4C4C4" rx="8" />
                                <path
                                    fill="#737373"
                                    d="M74.111 3.412A8 8 0 0180.665 0H100a8 8 0 018 8v24a8 8 0 01-8 8H48.5L74.111 3.412z"
                                />
                                <rect width="32" height="16" fill="#C4C4C4" rx="8" />
                                <rect width="32" height="16" y="24" fill="#737373" rx="8" />
                                <rect width="32" height="16" y="48" fill="#C4C4C4" rx="8" />
                            </svg>
                        </div>
                    </SelectableCard>
                </div>

                <h3 className="mt-12">
                    Dotfiles{" "}
                    <PillLabel type="warn" className="font-semibold mt-2 py-0.5 px-2 self-center">
                        Beta
                    </PillLabel>
                </h3>
                <p className="text-base text-gray-500 dark:text-gray-400">Customize workspaces using dotfiles.</p>
                <div className="mt-4 max-w-md">
                    <h4>Repository URL</h4>
                    <input
                        type="text"
                        value={dotfileRepo}
                        className="w-full"
                        placeholder="e.g. https://github.com/username/dotfiles"
                        onChange={(e) => setDotfileRepo(e.target.value)}
                    />
                    <div className="mt-1">
                        <p className="text-gray-500 dark:text-gray-400">
                            Add a repository URL that includes dotfiles. Gitpod will clone and install your dotfiles for
                            every new workspace.
                        </p>
                    </div>
                    <div className="mt-4 max-w-md">
                        <button onClick={() => actuallySetDotfileRepo(dotfileRepo)}>Save Changes</button>
                    </div>
                </div>
            </PageWithSubMenu>
        </div>
    );
}

function orderedIdeOptions(ideOptions: IDEOptions) {
    // TODO: Maybe convert orderKey to number before sort?
    return Object.entries(ideOptions.options)
        .filter(([_, x]) => !x.hidden)
        .sort((a, b) => {
            const keyA = a[1].orderKey || a[0];
            const keyB = b[1].orderKey || b[0];
            return keyA.localeCompare(keyB);
        });
}

function renderIdeOption(option: IDEOption, selected: boolean, onSelect: () => void): JSX.Element {
    const card = (
        <SelectableCardSolid className="w-36 h-40" title={option.title} selected={selected} onClick={onSelect}>
            <div className="flex justify-center mt-3">
                <img className="w-16 filter-grayscale self-center" src={option.logo} alt="logo" />
            </div>
            {option.label ? (
                <div
                    className={`font-semibold text-sm ${
                        selected ? "text-gray-100 dark:text-gray-600" : "text-gray-600 dark:text-gray-500"
                    } uppercase mt-2 px-3 py-1 self-center`}
                >
                    {option.label}
                </div>
            ) : (
                <></>
            )}
        </SelectableCardSolid>
    );

    if (option.tooltip) {
        return <Tooltip content={option.tooltip}>{card}</Tooltip>;
    }
    return card;
}
