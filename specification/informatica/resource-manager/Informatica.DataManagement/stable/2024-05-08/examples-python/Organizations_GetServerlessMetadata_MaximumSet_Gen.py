from azure.identity import DefaultAzureCredential

from azure.mgmt.informaticadatamanagement import InformaticaDataMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-informaticadatamanagement
# USAGE
    python organizations_get_serverless_metadata_maximum_set_gen.py

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

    response = client.organizations.get_serverless_metadata(
        resource_group_name="rgopenapi",
        organization_name="3_UC",
    )
    print(response)


# x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_GetServerlessMetadata_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
