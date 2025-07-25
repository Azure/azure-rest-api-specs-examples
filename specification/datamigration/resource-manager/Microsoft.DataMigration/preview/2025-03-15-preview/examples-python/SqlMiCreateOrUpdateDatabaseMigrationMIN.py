from azure.identity import DefaultAzureCredential

from azure.mgmt.datamigration import DataMigrationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datamigration
# USAGE
    python sql_mi_create_or_update_database_migration_min.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataMigrationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.database_migrations_sql_mi.begin_create_or_update(
        resource_group_name="testrg",
        managed_instance_name="managedInstance1",
        target_db_name="db1",
        parameters={
            "properties": {
                "backupConfiguration": {
                    "sourceLocation": {
                        "fileShare": {"password": "placeholder", "path": "C:\\aaa\\bbb\\ccc", "username": "name"}
                    },
                    "targetLocation": {
                        "accountKey": "abcd",
                        "storageAccountResourceId": "account.database.windows.net",
                    },
                },
                "kind": "SqlMi",
                "migrationService": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent",
                "scope": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance",
                "sourceDatabaseName": "aaa",
                "sourceSqlConnection": {
                    "authentication": "WindowsAuthentication",
                    "dataSource": "aaa",
                    "encryptConnection": True,
                    "password": "placeholder",
                    "trustServerCertificate": True,
                    "userName": "bbb",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlMiCreateOrUpdateDatabaseMigrationMIN.json
if __name__ == "__main__":
    main()
