/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import React, { createContext, useState } from "react";
import { GetLicenseInfoResult } from "@gitpod/gitpod-protocol";

const LicenseContext = createContext<{
    licenseSettings?: GetLicenseInfoResult;
    setLicenseSettings: React.Dispatch<GetLicenseInfoResult>;
}>({
    setLicenseSettings: () => null,
});

const LicenseContextProvider: React.FC = ({ children }) => {
    const [licenseSettings, setLicenseSettings] = useState<GetLicenseInfoResult>();
    return (
        <LicenseContext.Provider value={{ licenseSettings, setLicenseSettings }}>{children}</LicenseContext.Provider>
    );
};

export { LicenseContext, LicenseContextProvider };
