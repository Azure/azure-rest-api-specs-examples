import com.azure.core.util.Context;

/** Samples for CommunityGalleryImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/communityGalleryExamples/CommunityGalleryImage_Get.json
     */
    /**
     * Sample code: Get a community gallery image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACommunityGalleryImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCommunityGalleryImages()
            .getWithResponse("myLocation", "publicGalleryName", "myGalleryImageName", Context.NONE);
    }
}
