from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdbforpostgresql import CosmosdbForPostgresqlMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdbforpostgresql
# USAGE
    python configuration_update_coordinator.py

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

    response = client.configurations.begin_update_on_coordinator(
        resource_group_name="TestResourceGroup",
        cluster_name="testcluster",
        configuration_name="array_nulls",
        parameters={"properties": {"value": "on"}},
    ).result()
    print(response)


# x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/ConfigurationUpdateCoordinator.json
if __name__ == "__main__":
    main()
