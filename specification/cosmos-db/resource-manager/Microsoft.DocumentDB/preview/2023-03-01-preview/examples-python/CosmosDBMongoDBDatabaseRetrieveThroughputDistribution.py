from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_mongo_db_database_retrieve_throughput_distribution.py

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

    response = client.mongo_db_resources.begin_mongo_db_database_retrieve_throughput_distribution(
        resource_group_name="rg1",
        account_name="ddb1",
        database_name="databaseName",
        retrieve_throughput_parameters={
            "properties": {"resource": {"physicalPartitionIds": [{"id": "0"}, {"id": "1"}]}}
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBMongoDBDatabaseRetrieveThroughputDistribution.json
if __name__ == "__main__":
    main()
