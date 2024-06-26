from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python create_sql_pool_metadata_sync_config.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="01234567-89ab-4def-0123-456789abcdef",
    )

    response = client.sql_pool_metadata_sync_configs.create(
        resource_group_name="ExampleResourceGroup",
        workspace_name="ExampleWorkspace",
        sql_pool_name="ExampleSqlPool",
        metadata_sync_configuration={"properties": {"enabled": True}},
    )
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPoolMetadataSyncConfig.json
if __name__ == "__main__":
    main()
