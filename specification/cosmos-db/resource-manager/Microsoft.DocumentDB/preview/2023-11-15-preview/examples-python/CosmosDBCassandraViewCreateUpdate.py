from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_cassandra_view_create_update.py

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

    response = client.cassandra_resources.begin_create_update_cassandra_view(
        resource_group_name="rg1",
        account_name="ddb1",
        keyspace_name="keyspacename",
        view_name="viewname",
        create_update_cassandra_view_parameters={
            "properties": {
                "options": {},
                "resource": {
                    "id": "viewname",
                    "viewDefinition": "SELECT columna, columnb, columnc FROM keyspacename.srctablename WHERE columna IS NOT NULL AND columnc IS NOT NULL PRIMARY (columnc, columna)",
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBCassandraViewCreateUpdate.json
if __name__ == "__main__":
    main()
