from azure.identity import DefaultAzureCredential
from azure.mgmt.alertsmanagement import AlertsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-alertsmanagement
# USAGE
    python alert_processing_rules_create_or_update_remove_all_action_groups_outside_business_hours.py

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
        alert_processing_rule_name="RemoveActionGroupsOutsideBusinessHours",
        alert_processing_rule={
            "location": "Global",
            "properties": {
                "actions": [{"actionType": "RemoveAllActionGroups"}],
                "description": "Remove all ActionGroups outside business hours",
                "enabled": True,
                "schedule": {
                    "recurrences": [
                        {"endTime": "09:00:00", "recurrenceType": "Daily", "startTime": "17:00:00"},
                        {"daysOfWeek": ["Saturday", "Sunday"], "recurrenceType": "Weekly"},
                    ],
                    "timeZone": "Eastern Standard Time",
                },
                "scopes": ["/subscriptions/subId1"],
            },
            "tags": {},
        },
    )
    print(response)


# x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_remove_all_action_groups_outside_business_hours.json
if __name__ == "__main__":
    main()
