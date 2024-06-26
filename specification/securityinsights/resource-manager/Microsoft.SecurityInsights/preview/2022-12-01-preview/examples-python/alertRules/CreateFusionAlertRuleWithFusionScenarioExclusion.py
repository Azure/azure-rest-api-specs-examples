from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_fusion_alert_rule_with_fusion_scenario_exclusion.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.alert_rules.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        rule_id="myFirstFusionRule",
        alert_rule={
            "etag": "3d00c3ca-0000-0100-0000-5d42d5010000",
            "kind": "Fusion",
            "properties": {
                "alertRuleTemplateName": "f71aba3d-28fb-450b-b192-4e76a83015c8",
                "enabled": True,
                "sourceSettings": [
                    {"enabled": True, "sourceName": "Anomalies", "sourceSubTypes": None},
                    {
                        "enabled": True,
                        "sourceName": "Alert providers",
                        "sourceSubTypes": [
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Azure Active Directory Identity Protection",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Azure Defender",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Azure Defender for IoT",
                            },
                            {
                                "enabled": True,
                                "severityFilter": ["High", "Medium", "Low", "Informational"],
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Microsoft 365 Defender",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Microsoft Cloud App Security",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Microsoft Defender for Endpoint",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Microsoft Defender for Identity",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Microsoft Defender for Office 365",
                            },
                            {
                                "enabled": True,
                                "severityFilters": {
                                    "filters": [
                                        {"enabled": True, "severity": "High"},
                                        {"enabled": True, "severity": "Medium"},
                                        {"enabled": True, "severity": "Low"},
                                        {"enabled": True, "severity": "Informational"},
                                    ]
                                },
                                "sourceSubTypeName": "Azure Sentinel scheduled analytics rules",
                            },
                        ],
                    },
                    {
                        "enabled": True,
                        "sourceName": "Raw logs from other sources",
                        "sourceSubTypes": [
                            {
                                "enabled": True,
                                "severityFilters": {"filters": None},
                                "sourceSubTypeName": "Palo Alto Networks",
                            }
                        ],
                    },
                ],
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
if __name__ == "__main__":
    main()
