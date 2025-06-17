from azure.identity import DefaultAzureCredential

from azure.mgmt.datamigration import DataMigrationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datamigration
# USAGE
    python sql_db_cancel_database_migration.py

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

    client.database_migrations_sql_db.begin_cancel(
        resource_group_name="testrg",
        sql_db_instance_name="sqldbinstance",
        target_db_name="db1",
        parameters={"migrationOperationId": "9a90bb84-e70f-46f7-b0ae-1aef5b3b9f07"},
    ).result()


# x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlDbCancelDatabaseMigration.json
if __name__ == "__main__":
    main()
