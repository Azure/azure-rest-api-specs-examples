from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_sql_container_redistribute_throughput.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.sql_resources.begin_sql_container_redistribute_throughput(
        resource_group_name="rg1",
        account_name="ddb1",
        database_name="databaseName",
        container_name="containerName",
        redistribute_throughput_parameters={
            "properties": {
                "resource": {
                    "sourcePhysicalPartitionThroughputInfo": [{"id": "2", "throughput": 5000}, {"id": "3"}],
                    "targetPhysicalPartitionThroughputInfo": [
                        {"id": "0", "throughput": 5000},
                        {"id": "1", "throughput": 5000},
                    ],
                    "throughputPolicy": "custom",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBSqlContainerRedistributeThroughput.json
if __name__ == "__main__":
    main()
