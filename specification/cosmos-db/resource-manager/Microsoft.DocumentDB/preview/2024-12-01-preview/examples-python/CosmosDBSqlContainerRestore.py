from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_sql_container_restore.py

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

    response = client.sql_resources.begin_create_update_sql_container(
        resource_group_name="rg1",
        account_name="ddb1",
        database_name="databaseName",
        container_name="containerName",
        create_update_sql_container_parameters={
            "location": "West US",
            "properties": {
                "options": {},
                "resource": {
                    "createMode": "Restore",
                    "id": "containerName",
                    "restoreParameters": {
                        "restoreSource": "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId",
                        "restoreTimestampInUtc": "2022-07-20T18:28:00Z",
                        "restoreWithTtlDisabled": True,
                    },
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBSqlContainerRestore.json
if __name__ == "__main__":
    main()
