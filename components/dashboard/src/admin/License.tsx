/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { PageWithSubMenu } from "../components/PageWithSubMenu";
import { adminMenu } from "./admin-menu";

import { LicenseContext } from "../license-context";
import { useContext } from "react";
import { getGitpodService } from "../service/service";
import { GetLicenseInfoResult } from "@gitpod/gitpod-protocol";

export default function License() {
    const { licenseSettings, setLicenseSettings } = useContext(LicenseContext);

    // @ts-ignore
    const actuallySetLicenseSettings = async (value: GetLicenseInfoResult) => {
        await getGitpodService().server.getLicenseInfo();
        setLicenseSettings(value);
    };

    return (
        <div>
            <PageWithSubMenu subMenu={adminMenu} title="License" subtitle="License information of your account.">
                <>
                    This info is about the license
                    <h3> {licenseSettings?.isAdmin} </h3>
                </>
            </PageWithSubMenu>
        </div>
    );
}
