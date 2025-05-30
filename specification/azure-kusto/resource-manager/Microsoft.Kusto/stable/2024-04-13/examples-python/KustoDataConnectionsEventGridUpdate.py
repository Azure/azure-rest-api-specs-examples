from azure.identity import DefaultAzureCredential

from azure.mgmt.kusto import KustoManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kusto
# USAGE
    python kusto_data_connections_event_grid_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KustoManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-123456789098",
    )

    response = client.data_connections.begin_update(
        resource_group_name="kustorptest",
        cluster_name="kustoCluster",
        database_name="KustoDatabase8",
        data_connection_name="dataConnectionTest",
        parameters={
            "kind": "EventGrid",
            "location": "westus",
            "properties": {
                "blobStorageEventType": "Microsoft.Storage.BlobCreated",
                "consumerGroup": "$Default",
                "dataFormat": "MULTIJSON",
                "databaseRouting": "Single",
                "eventGridResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest",
                "eventHubResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest2",
                "ignoreFirstRecord": False,
                "managedIdentityResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
                "mappingRuleName": "TestMapping",
                "storageAccountResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount",
                "tableName": "TestTable",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDataConnectionsEventGridUpdate.json
if __name__ == "__main__":
    main()
