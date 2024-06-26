from azure.identity import DefaultAzureCredential
from azure.mgmt.databricks import AzureDatabricksManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databricks
# USAGE
    python private_link_resources_get.py

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

    response = client.private_link_resources.get(
        resource_group_name="myResourceGroup",
        workspace_name="myWorkspace",
        group_id="databricks_ui_api",
    )
    print(response)


# x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateLinkResourcesGet.json
if __name__ == "__main__":
    main()
