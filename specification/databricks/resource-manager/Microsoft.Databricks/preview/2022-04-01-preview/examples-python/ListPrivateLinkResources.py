from azure.identity import DefaultAzureCredential
from azure.mgmt.databricks import AzureDatabricksManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databricks
# USAGE
    python list_private_link_resources.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureDatabricksManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="11111111-1111-1111-1111-111111111111",
    )

    response = client.private_link_resources.list(
        resource_group_name="myResourceGroup",
        workspace_name="myWorkspace",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2022-04-01-preview/examples/ListPrivateLinkResources.json
if __name__ == "__main__":
    main()
