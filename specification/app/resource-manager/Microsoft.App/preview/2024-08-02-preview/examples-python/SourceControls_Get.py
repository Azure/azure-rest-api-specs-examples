from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python source_controls_get.py

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

    response = client.container_apps_source_controls.get(
        resource_group_name="workerapps-rg-xj",
        container_app_name="testcanadacentral",
        source_control_name="current",
    )
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/SourceControls_Get.json
if __name__ == "__main__":
    main()
