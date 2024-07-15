from azure.identity import DefaultAzureCredential

from azure.mgmt.informaticadatamanagement import InformaticaDataMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-informaticadatamanagement
# USAGE
    python organizations_list_by_resource_group_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = InformaticaDataMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="3599DA28-E346-4D9F-811E-189C0445F0FE",
    )

    response = client.organizations.list_by_resource_group(
        resource_group_name="rgopenapi",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_ListByResourceGroup_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
