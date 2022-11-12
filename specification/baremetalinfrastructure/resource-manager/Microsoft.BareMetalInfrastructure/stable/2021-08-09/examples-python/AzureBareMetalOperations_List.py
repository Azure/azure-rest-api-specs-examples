from azure.identity import DefaultAzureCredential
from azure.mgmt.baremetalinfrastructure import BareMetalInfrastructureClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-baremetalinfrastructure
# USAGE
    python list_all_management_operations_supported_by_the_azure_bare_metal_rp.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BareMetalInfrastructureClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalOperations_List.json
if __name__ == "__main__":
    main()
