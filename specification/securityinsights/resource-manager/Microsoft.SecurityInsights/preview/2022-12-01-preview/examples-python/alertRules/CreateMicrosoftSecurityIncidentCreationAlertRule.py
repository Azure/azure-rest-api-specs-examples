from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_microsoft_security_incident_creation_alert_rule.py

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
        rule_id="microsoftSecurityIncidentCreationRuleExample",
        alert_rule={
            "etag": '"260097e0-0000-0d00-0000-5d6fa88f0000"',
            "kind": "MicrosoftSecurityIncidentCreation",
            "properties": {
                "displayName": "testing displayname",
                "enabled": True,
                "productFilter": "Microsoft Cloud App Security",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/alertRules/CreateMicrosoftSecurityIncidentCreationAlertRule.json
if __name__ == "__main__":
    main()
