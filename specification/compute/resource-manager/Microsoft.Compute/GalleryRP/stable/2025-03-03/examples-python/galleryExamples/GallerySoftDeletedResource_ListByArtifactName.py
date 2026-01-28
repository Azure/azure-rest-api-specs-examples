from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python gallery_soft_deleted_resource_list_by_artifact_name.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.soft_deleted_resource.list_by_artifact_name(
        resource_group_name="myResourceGroup",
        gallery_name="myGalleryName",
        artifact_type="images",
        artifact_name="myGalleryImageName",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GallerySoftDeletedResource_ListByArtifactName.json
if __name__ == "__main__":
    main()
