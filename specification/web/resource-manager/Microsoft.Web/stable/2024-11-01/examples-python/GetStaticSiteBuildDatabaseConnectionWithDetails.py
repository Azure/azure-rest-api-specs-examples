from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python get_static_site_build_database_connection_with_details.py

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

    response = client.static_sites.get_build_database_connection_with_details(
        resource_group_name="rg",
        name="testStaticSite0",
        environment_name="default",
        database_connection_name="default",
    )
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetStaticSiteBuildDatabaseConnectionWithDetails.json
if __name__ == "__main__":
    main()
