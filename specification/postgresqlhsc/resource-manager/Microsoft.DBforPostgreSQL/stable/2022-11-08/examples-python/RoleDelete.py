from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdbforpostgresql import CosmosdbForPostgresqlMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdbforpostgresql
# USAGE
    python role_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosdbForPostgresqlMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    client.roles.begin_delete(
        resource_group_name="TestGroup",
        cluster_name="pgtestsvc4",
        role_name="role1",
    ).result()


# x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/RoleDelete.json
if __name__ == "__main__":
    main()
