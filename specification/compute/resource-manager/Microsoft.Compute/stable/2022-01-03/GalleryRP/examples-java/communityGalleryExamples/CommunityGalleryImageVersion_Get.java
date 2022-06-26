import com.azure.core.util.Context;

/** Samples for CommunityGalleryImageVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/communityGalleryExamples/CommunityGalleryImageVersion_Get.json
     */
    /**
     * Sample code: Get a community gallery image version.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACommunityGalleryImageVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCommunityGalleryImageVersions()
            .getWithResponse(
                "myLocation", "publicGalleryName", "myGalleryImageName", "myGalleryImageVersionName", Context.NONE);
    }
}
