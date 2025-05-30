from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python container_apps_diagnostics_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerAppsAPIClient(
        credential=DefaultAzureCredential(),
        subscription_id="f07f3711-b45e-40fe-a941-4e6d93f851e6",
    )

    response = client.container_apps_diagnostics.list_detectors(
        resource_group_name="mikono-workerapp-test-rg",
        container_app_name="mikono-capp-stage1",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsDiagnostics_List.json
if __name__ == "__main__":
    main()
