from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.sphere import AzureSphereMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sphere
# USAGE
    python post_upload_image_catalog.py

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

    client.catalogs.begin_upload_image(
        resource_group_name="MyResourceGroup1",
        catalog_name="MyCatalog1",
        upload_image_request={"properties": {"image": "bXliYXNlNjRzdHJpbmc="}},
    ).result()


# x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostUploadImageCatalog.json
if __name__ == "__main__":
    main()
