from azure.identity import DefaultAzureCredential

from azure.mgmt.kusto import KustoManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kusto
# USAGE
    python kusto_data_connection_validation_async.py

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

    response = client.data_connections.begin_data_connection_validation(
        resource_group_name="kustorptest",
        cluster_name="kustoCluster",
        database_name="KustoDatabase8",
        parameters={
            "dataConnectionName": "dataConnectionTest",
            "properties": {
                "kind": "EventHub",
                "properties": {
                    "compression": "None",
                    "consumerGroup": "testConsumerGroup1",
                    "dataFormat": "MULTIJSON",
                    "eventHubResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1",
                    "managedIdentityResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
                    "mappingRuleName": "TestMapping",
                    "tableName": "TestTable",
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDataConnectionValidationAsync.json
if __name__ == "__main__":
    main()
