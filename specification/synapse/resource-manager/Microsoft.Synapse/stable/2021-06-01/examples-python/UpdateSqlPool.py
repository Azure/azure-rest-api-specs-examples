from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python update_sql_pool.py

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

    response = client.sql_pools.begin_update(
        resource_group_name="ExampleResourceGroup",
        workspace_name="ExampleWorkspace",
        sql_pool_name="ExampleSqlPool",
        sql_pool_info={
            "location": "West US 2",
            "properties": {"collation": "", "maxSizeBytes": 0, "restorePointInTime": "1970-01-01T00:00:00.000Z"},
            "sku": {"name": "", "tier": ""},
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateSqlPool.json
if __name__ == "__main__":
    main()
