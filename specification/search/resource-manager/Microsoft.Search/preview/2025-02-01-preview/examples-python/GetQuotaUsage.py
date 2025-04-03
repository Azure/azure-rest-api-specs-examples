from azure.identity import DefaultAzureCredential

from azure.mgmt.search import SearchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-search
# USAGE
    python get_quota_usage.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SearchManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.usage_by_subscription_sku(
        location="westus",
        sku_name="free",
    )
    print(response)


# x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/GetQuotaUsage.json
if __name__ == "__main__":
    main()
