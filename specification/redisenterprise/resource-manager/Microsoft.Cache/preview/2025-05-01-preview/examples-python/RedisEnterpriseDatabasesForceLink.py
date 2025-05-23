from azure.identity import DefaultAzureCredential

from azure.mgmt.redisenterprise import RedisEnterpriseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-redisenterprise
# USAGE
    python redis_enterprise_databases_force_link.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RedisEnterpriseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.databases.begin_force_link_to_replication_group(
        resource_group_name="rg1",
        cluster_name="cache1",
        database_name="default",
        parameters={
            "geoReplication": {
                "groupNickname": "groupName",
                "linkedDatabases": [
                    {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"
                    },
                    {
                        "id": "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"
                    },
                ],
            }
        },
    ).result()


# x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/RedisEnterpriseDatabasesForceLink.json
if __name__ == "__main__":
    main()
