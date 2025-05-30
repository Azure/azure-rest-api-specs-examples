from azure.identity import DefaultAzureCredential

from azure.mgmt.containerinstance import ContainerInstanceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerinstance
# USAGE
    python container_groups_start.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerInstanceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.container_groups.begin_start(
        resource_group_name="demo",
        container_group_name="demo1",
    ).result()


# x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupsStart.json
if __name__ == "__main__":
    main()
