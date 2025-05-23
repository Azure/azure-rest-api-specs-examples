from azure.identity import DefaultAzureCredential

from azure.mgmt.kusto import KustoManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-kusto
# USAGE
    python kusto_data_connections_cosmos_db_update.py

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
        database_name="KustoDatabase1",
        data_connection_name="dataConnectionTest",
        parameters={
            "kind": "CosmosDb",
            "location": "westus",
            "properties": {
                "cosmosDbAccountResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1",
                "cosmosDbContainer": "cosmosDbContainerTest",
                "cosmosDbDatabase": "cosmosDbDatabaseTest",
                "managedIdentityResourceId": "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
                "mappingRuleName": "TestMapping",
                "retrievalStartDate": "2022-07-29T12:00:00.6554616Z",
                "tableName": "TestTable",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDataConnectionsCosmosDbUpdate.json
if __name__ == "__main__":
    main()
