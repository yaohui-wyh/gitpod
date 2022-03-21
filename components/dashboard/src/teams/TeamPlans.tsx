/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { TeamMemberInfo } from "@gitpod/gitpod-protocol";
import { Currency, Plans } from "@gitpod/gitpod-protocol/lib/plans";
import { TeamSubscription2 } from "@gitpod/gitpod-protocol/lib/team-subscription-protocol";
import { useContext, useEffect, useState } from "react";
import { useLocation } from "react-router";
import { ChargebeeClient } from "../chargebee/chargebee-client";
import { PageWithSubMenu } from "../components/PageWithSubMenu";
import SelectableCard from "../components/SelectableCard";
import { PaymentContext } from "../payment-context";
import { getGitpodService } from "../service/service";
import { getCurrentTeam, TeamsContext } from "./teams-context";
import { getTeamSettingsMenu } from "./TeamSettings";

export default function TeamPlans() {
    const { teams } = useContext(TeamsContext);
    const location = useLocation();
    const team = getCurrentTeam(location, teams);
    const [members, setMembers] = useState<TeamMemberInfo[]>([]);
    const [teamSubscription, setTeamSubscription] = useState<TeamSubscription2 | undefined>();
    const { showPaymentUI, currency, setCurrency } = useContext(PaymentContext);

    useEffect(() => {
        if (!team) {
            return;
        }
        (async () => {
            const [memberInfos, subscription] = await Promise.all([
                getGitpodService().server.getTeamMembers(team.id),
                getGitpodService().server.getTeamSubscription(team.id),
            ]);
            setMembers(memberInfos);
            setTeamSubscription(subscription);
        })();
    }, [team]);

    const availableTeamPlans = Plans.getAvailableTeamPlans(currency || "USD");

    const checkout = async (chargebeePlanId: string) => {
        if (!team || members.length < 1) {
            return;
        }
        const chargebeeClient = await ChargebeeClient.getOrCreate();
        await new Promise((resolve, reject) => {
            chargebeeClient.checkout((paymentServer) => paymentServer.teamCheckout(team.id, chargebeePlanId), {
                success: resolve,
                error: reject,
            });
        });
    };

    return (
        <PageWithSubMenu
            subMenu={getTeamSettingsMenu({ team, showPaymentUI })}
            title="Plans"
            subtitle="Manage team plans and billing."
        >
            <p className="text-sm">
                <a
                    className={`gp-link ${isChargebeeCustomer ? "" : "invisible"}`}
                    href="javascript:void(0)"
                    onClick={() => {
                        ChargebeeClient.getOrCreate().then((chargebeeClient) => chargebeeClient.openPortal());
                    }}
                >
                    Billing
                </a>
                <span className="pl-6">
                    {currency === "EUR" ? (
                        <>
                            € /{" "}
                            <a className="gp-link" href="javascript:void(0)" onClick={() => setCurrency("USD")}>
                                $
                            </a>
                        </>
                    ) : (
                        <>
                            <a className="gp-link" href="javascript:void(0)" onClick={() => setCurrency("EUR")}>
                                €
                            </a>{" "}
                            / $
                        </>
                    )}
                </span>
            </p>
            <div className="mt-4 space-x-4 flex">
                <SelectableCard
                    className="w-36 h-32"
                    title="Free"
                    selected={!teamSubscription?.planId}
                    onClick={() => {}}
                >
                    {members.length} x {Currency.getSymbol(currency || "USD")}0 ={" "}
                    {Currency.getSymbol(currency || "USD")}0
                </SelectableCard>
                {availableTeamPlans.map((tp) => (
                    <SelectableCard
                        className="w-36 h-32"
                        title={tp.name}
                        selected={tp.chargebeeId === teamSubscription?.planId}
                        onClick={() => checkout(tp.chargebeeId)}
                    >
                        {members.length} x {Currency.getSymbol(tp.currency)}
                        {tp.pricePerMonth} = {Currency.getSymbol(tp.currency)}
                        {members.length * tp.pricePerMonth}
                    </SelectableCard>
                ))}
            </div>
        </PageWithSubMenu>
    );
}
