from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python get_subscription_operation_with_async_response.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    client.global_operations.get_subscription_operation_with_async_response(
        location="West US",
        operation_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab5d5",
    )


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetSubscriptionOperationWithAsyncResponse.json
if __name__ == "__main__":
    main()
