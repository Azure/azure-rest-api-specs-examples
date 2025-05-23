from azure.identity import DefaultAzureCredential

from azure.mgmt.databasewatcher import DatabaseWatcherMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databasewatcher
# USAGE
    python alert_rule_resources_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DatabaseWatcherMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.alert_rule_resources.create_or_update(
        resource_group_name="rgWatcher",
        watcher_name="testWatcher",
        alert_rule_resource_name="testAlert",
        resource={
            "properties": {
                "alertRuleResourceId": "/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878Be/resourceGroups/rgWatcher/providers/microsoft.insights/scheduledqueryrules/alerts-demo",
                "alertRuleTemplateId": "someTemplateId",
                "alertRuleTemplateVersion": "1.0",
                "createdWithProperties": "CreatedWithActionGroup",
                "creationTime": "2024-07-25T15:38:47.798Z",
            }
        },
    )
    print(response)


# x-ms-original-file: 2025-01-02/AlertRuleResources_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
