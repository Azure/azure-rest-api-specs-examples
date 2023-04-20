from azure.identity import DefaultAzureCredential
from azure.mgmt.redisenterprise import RedisEnterpriseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-redisenterprise
# USAGE
    python redis_enterprise_skus_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RedisEnterpriseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.skus.list(
        location="westus2",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseSkusList.json
if __name__ == "__main__":
    main()
