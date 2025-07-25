from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python update_auth_settings.py

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

    response = client.web_apps.update_auth_settings(
        resource_group_name="testrg123",
        name="sitef6141",
        site_auth_settings={
            "properties": {
                "allowedExternalRedirectUrls": ["sitef6141.customdomain.net", "sitef6141.customdomain.info"],
                "clientId": "42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com",
                "defaultProvider": "Google",
                "enabled": True,
                "runtimeVersion": "~1",
                "tokenRefreshExtensionHours": 120,
                "tokenStoreEnabled": True,
                "unauthenticatedClientAction": "RedirectToLoginPage",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateAuthSettings.json
if __name__ == "__main__":
    main()
