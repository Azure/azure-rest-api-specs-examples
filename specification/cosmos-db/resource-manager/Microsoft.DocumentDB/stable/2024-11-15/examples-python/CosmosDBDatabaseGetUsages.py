from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_database_get_usages.py

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

    response = client.database.list_usages(
        resource_group_name="rg1",
        account_name="ddb1",
        database_rid="databaseRid",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBDatabaseGetUsages.json
if __name__ == "__main__":
    main()
