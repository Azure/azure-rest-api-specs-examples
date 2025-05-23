from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python container_apps_get_auth_token.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerAppsAPIClient(
        credential=DefaultAzureCredential(),
        subscription_id="651f8027-33e8-4ec4-97b4-f6e9f3dc8744",
    )

    response = client.container_apps.get_auth_token(
        resource_group_name="rg",
        container_app_name="testcontainerApp0",
    )
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerApps_GetAuthToken.json
if __name__ == "__main__":
    main()
