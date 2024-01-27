from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_throughput_pool_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.throughput_pool.begin_create_or_update(
        resource_group_name="rg1",
        throughput_pool_name="tp1",
        body={"location": "westus2", "properties": {"maxThroughput": 10000}, "tags": {}},
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/throughputPool/CosmosDBThroughputPoolCreate.json
if __name__ == "__main__":
    main()
