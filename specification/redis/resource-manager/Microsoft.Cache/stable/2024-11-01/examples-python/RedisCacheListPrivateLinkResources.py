from azure.identity import DefaultAzureCredential

from azure.mgmt.redis import RedisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-redis
# USAGE
    python redis_cache_list_private_link_resources.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RedisManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.private_link_resources.list_by_redis_cache(
        resource_group_name="rgtest01",
        cache_name="cacheTest01",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheListPrivateLinkResources.json
if __name__ == "__main__":
    main()
