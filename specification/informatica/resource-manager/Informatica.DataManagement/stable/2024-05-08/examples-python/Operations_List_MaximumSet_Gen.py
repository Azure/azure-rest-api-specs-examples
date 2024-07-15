from azure.identity import DefaultAzureCredential

from azure.mgmt.informaticadatamanagement import InformaticaDataMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-informaticadatamanagement
# USAGE
    python operations_list_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = InformaticaDataMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Operations_List_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
