from azure.identity import DefaultAzureCredential
from azure.mgmt.azurearcdata import AzureArcDataManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurearcdata
# USAGE
    python update_data_controller.py

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

    response = client.data_controllers.begin_patch_data_controller(
        resource_group_name="testrg",
        data_controller_name="testdataController1",
        data_controller_resource={"tags": {"mytag": "myval"}},
    ).result()
    print(response)


# x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/UpdateDataController.json
if __name__ == "__main__":
    main()
