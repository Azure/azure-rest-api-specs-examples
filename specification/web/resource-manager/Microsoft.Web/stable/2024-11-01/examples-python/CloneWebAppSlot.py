from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python clone_web_app_slot.py

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

    response = client.web_apps.begin_create_or_update_slot(
        resource_group_name="testrg123",
        name="sitef6141",
        slot="staging",
        site_envelope={
            "kind": "app",
            "location": "East US",
            "properties": {
                "cloningInfo": {
                    "appSettingsOverrides": {"Setting1": "NewValue1", "Setting3": "NewValue5"},
                    "cloneCustomHostNames": True,
                    "cloneSourceControl": True,
                    "configureLoadBalancing": False,
                    "hostingEnvironment": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/hostingenvironments/aseforsites",
                    "overwrite": False,
                    "sourceWebAppId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/sites/srcsiteg478/slot/qa",
                    "sourceWebAppLocation": "West Europe",
                }
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/CloneWebAppSlot.json
if __name__ == "__main__":
    main()
