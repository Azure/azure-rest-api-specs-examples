from azure.identity import DefaultAzureCredential
from azure.mgmt.azurearcdata import AzureArcDataManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurearcdata
# USAGE
    python get_active_directory_connector.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureArcDataManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.active_directory_connectors.get(
        resource_group_name="testrg",
        data_controller_name="testdataController",
        active_directory_connector_name="testADConnector",
    )
    print(response)


# x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/GetActiveDirectoryConnector.json
if __name__ == "__main__":
    main()
