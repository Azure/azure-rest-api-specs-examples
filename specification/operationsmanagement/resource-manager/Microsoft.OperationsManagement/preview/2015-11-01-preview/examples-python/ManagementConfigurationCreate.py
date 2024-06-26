from azure.identity import DefaultAzureCredential
from azure.mgmt.operationsmanagement import OperationsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-operationsmanagement
# USAGE
    python management_configuration_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OperationsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.management_configurations.create_or_update(
        resource_group_name="rg1",
        management_configuration_name="managementConfiguration1",
        parameters={"location": "East US"},
    )
    print(response)


# x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementConfigurationCreate.json
if __name__ == "__main__":
    main()
