from azure.identity import DefaultAzureCredential

from azure.mgmt.mongodbatlas import MongoDBAtlasMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mongodbatlas
# USAGE
    python organizations_delete_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MongoDBAtlasMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.organizations.begin_delete(
        resource_group_name="rgopenapi",
        organization_name="U.1-:7",
    ).result()


# x-ms-original-file: 2025-06-01/Organizations_Delete_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
