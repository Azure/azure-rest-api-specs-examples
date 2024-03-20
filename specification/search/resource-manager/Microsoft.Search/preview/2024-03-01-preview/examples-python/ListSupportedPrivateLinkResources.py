from azure.identity import DefaultAzureCredential
from azure.mgmt.search import SearchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-search
# USAGE
    python list_supported_private_link_resources.py

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

    response = client.private_link_resources.list_supported(
        resource_group_name="rg1",
        search_service_name="mysearchservice",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/ListSupportedPrivateLinkResources.json
if __name__ == "__main__":
    main()
