from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_nrt_alert_rule.py

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
        rule_id="73e01a99-5cd7-4139-a149-9f2736ff2ab5",
        alert_rule={
            "etag": '"0300bf09-0000-0000-0000-5c37296e0000"',
            "kind": "NRT",
            "properties": {
                "description": "",
                "displayName": "Rule2",
                "enabled": True,
                "eventGroupingSettings": {"aggregationKind": "AlertPerResult"},
                "incidentConfiguration": {
                    "createIncident": True,
                    "groupingConfiguration": {
                        "enabled": True,
                        "groupByEntities": ["Host", "Account"],
                        "lookbackDuration": "PT5H",
                        "matchingMethod": "Selected",
                        "reopenClosedIncident": False,
                    },
                },
                "query": "ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity = ComputerIP_Hidden",
                "severity": "High",
                "suppressionDuration": "PT1H",
                "suppressionEnabled": False,
                "tactics": ["Persistence", "LateralMovement"],
                "techniques": ["T1037", "T1021"],
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/alertRules/CreateNrtAlertRule.json
if __name__ == "__main__":
    main()
