from azure.identity import DefaultAzureCredential

from azure.mgmt.postgresqlflexibleservers import PostgreSQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-postgresqlflexibleservers
# USAGE
    python migrations_get_migration_with_successful_validation_but_migration_failure.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PostgreSQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.migrations.get(
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
        resource_group_name="testrg",
        target_db_server_name="testtarget",
        migration_name="testmigrationwithsuccessfulvalidationbutmigrationfailure",
    )
    print(response)


# x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/Migrations_GetMigrationWithSuccessfulValidationButMigrationFailure.json
if __name__ == "__main__":
    main()
