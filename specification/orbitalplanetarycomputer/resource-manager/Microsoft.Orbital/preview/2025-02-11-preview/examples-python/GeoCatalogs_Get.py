from azure.identity import DefaultAzureCredential

from azure.mgmt.planetarycomputer import PlanetaryComputerMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-planetarycomputer
# USAGE
    python geo_catalogs_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PlanetaryComputerMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.geo_catalogs.get(
        resource_group_name="MyResourceGroup",
        catalog_name="MyCatalog",
    )
    print(response)


# x-ms-original-file: 2025-02-11-preview/GeoCatalogs_Get.json
if __name__ == "__main__":
    main()
