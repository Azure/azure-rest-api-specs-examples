from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python get_site_deployment_status_slot.py

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

    response = client.web_apps.begin_get_slot_site_deployment_status_slot(
        resource_group_name="rg",
        name="testSite",
        slot="stage",
        deployment_status_id="eacfd68b-3bbd-4ad9-99c5-98614d89c8e5",
    ).result()
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetSiteDeploymentStatusSlot.json
if __name__ == "__main__":
    main()
