from azure.identity import DefaultAzureCredential

from azure.mgmt.postgresqlflexibleservers import PostgreSQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-postgresqlflexibleservers
# USAGE
    python long_term_retention_operation_get.py

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

    response = client.ltr_backup_operations.get(
        resource_group_name="rgLongTermRetention",
        server_name="pgsqlltrtestserver",
        backup_name="backup1",
    )
    print(response)


# x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/LongTermRetentionOperationGet.json
if __name__ == "__main__":
    main()
