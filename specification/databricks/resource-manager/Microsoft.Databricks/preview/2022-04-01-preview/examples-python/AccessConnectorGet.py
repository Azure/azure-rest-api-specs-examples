from azure.identity import DefaultAzureCredential
from azure.mgmt.databricks import AzureDatabricksManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databricks
# USAGE
    python access_connector_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureDatabricksManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.access_connectors.get(
        resource_group_name="rg",
        connector_name="myAccessConnector",
    )
    print(response)


# x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2022-04-01-preview/examples/AccessConnectorGet.json
if __name__ == "__main__":
    main()
