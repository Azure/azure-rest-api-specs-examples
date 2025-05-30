from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python create_database_standby_mode.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.databases.begin_create_or_update(
        resource_group_name="Default-SQL-SouthEastAsia",
        server_name="testsvr",
        database_name="testdb",
        parameters={
            "location": "southeastasia",
            "properties": {
                "createMode": "Secondary",
                "secondaryType": "Standby",
                "sourceDatabaseId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Sql/servers/testsvr1/databases/testdb",
            },
            "sku": {"name": "S0", "tier": "Standard"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/CreateDatabaseStandbyMode.json
if __name__ == "__main__":
    main()
