from azure.identity import DefaultAzureCredential
from azure.mgmt.servicelinker import ServiceLinkerManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicelinker
# USAGE
    python configuration_names_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceLinkerManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.configuration_names.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/ConfigurationNamesList.json
if __name__ == "__main__":
    main()
