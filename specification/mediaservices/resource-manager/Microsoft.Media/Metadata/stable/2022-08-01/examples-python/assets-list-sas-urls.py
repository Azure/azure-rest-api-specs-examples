from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python assetslistsasurls.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureMediaServices(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.assets.list_container_sas(
        resource_group_name="contoso",
        account_name="contosomedia",
        asset_name="ClimbingMountBaker",
        parameters={"expiryTime": "2018-01-01T10:00:00.007Z", "permissions": "ReadWrite"},
    )
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-sas-urls.json
if __name__ == "__main__":
    main()
