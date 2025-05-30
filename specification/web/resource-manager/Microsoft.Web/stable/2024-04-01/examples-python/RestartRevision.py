from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python restart_revision.py

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

    client.container_apps_revisions.restart_revision(
        resource_group_name="rg",
        container_app_name="testStaticSite0",
        name="testcontainerApp0-pjxhsye",
    )


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/RestartRevision.json
if __name__ == "__main__":
    main()
