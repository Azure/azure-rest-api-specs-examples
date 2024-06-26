from azure.identity import DefaultAzureCredential
from azure.mgmt.alertsmanagement import AlertsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-alertsmanagement
# USAGE
    python alert_processing_rules_create_or_update_add_action_group_all_alerts_in_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AlertsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId1",
    )

    response = client.alert_processing_rules.create_or_update(
        resource_group_name="alertscorrelationrg",
        alert_processing_rule_name="AddActionGroupToSubscription",
        alert_processing_rule={
            "location": "Global",
            "properties": {
                "actions": [
                    {
                        "actionGroupIds": [
                            "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1"
                        ],
                        "actionType": "AddActionGroups",
                    }
                ],
                "description": "Add ActionGroup1 to all alerts in the subscription",
                "enabled": True,
                "scopes": ["/subscriptions/subId1"],
            },
            "tags": {},
        },
    )
    print(response)


# x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
if __name__ == "__main__":
    main()
