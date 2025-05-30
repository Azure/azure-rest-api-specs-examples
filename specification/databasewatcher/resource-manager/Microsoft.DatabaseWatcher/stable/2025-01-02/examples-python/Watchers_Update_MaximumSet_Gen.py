from azure.identity import DefaultAzureCredential

from azure.mgmt.databasewatcher import DatabaseWatcherMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databasewatcher
# USAGE
    python watchers_update_maximum_set_gen.py

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

    response = client.watchers.begin_update(
        resource_group_name="rgWatcher",
        watcher_name="testWatcher",
        properties={
            "identity": {"type": "SystemAssigned"},
            "properties": {
                "datastore": {
                    "adxClusterResourceId": "/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest/providers/Microsoft.Kusto/clusters/apiTestKusto",
                    "kustoClusterDisplayName": "kustoUri-adx",
                    "kustoClusterUri": "https://kustouri-adx.eastus.kusto.windows.net",
                    "kustoDataIngestionUri": "https://ingest-kustouri-adx.eastus.kusto.windows.net",
                    "kustoDatabaseName": "kustoDatabaseName1",
                    "kustoManagementUrl": "https://portal.azure.com/",
                    "kustoOfferingType": "adx",
                },
                "defaultAlertRuleIdentityResourceId": "/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878B/resourceGroups/rgWatcher/providers/Microsoft.ManagedIdentity/userAssignedIdentities/newtest",
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-01-02/Watchers_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
