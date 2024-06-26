from azure.identity import DefaultAzureCredential

from azure.mgmt.sphere import AzureSphereMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sphere
# USAGE
    python get_catalogs_rg.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureSphereMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.catalogs.list_by_resource_group(
        resource_group_name="MyResourceGroup1",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetCatalogsRG.json
if __name__ == "__main__":
    main()
