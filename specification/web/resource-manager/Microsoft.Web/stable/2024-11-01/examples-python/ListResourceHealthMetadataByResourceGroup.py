from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python list_resource_health_metadata_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="4adb32ad-8327-4cbb-b775-b84b4465bb38",
    )

    response = client.resource_health_metadata.list_by_resource_group(
        resource_group_name="Default-Web-NorthCentralUS",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/ListResourceHealthMetadataByResourceGroup.json
if __name__ == "__main__":
    main()
