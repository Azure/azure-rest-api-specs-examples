from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python builds_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerAppsAPIClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.builds.begin_delete(
        resource_group_name="rg",
        builder_name="testBuilder",
        build_name="testBuild",
    ).result()


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Builds_Delete.json
if __name__ == "__main__":
    main()
