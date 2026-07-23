
/**
 * Samples for CommunityGalleryImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/communityGalleryExamples/CommunityGalleryImage_Get.json
     */
    /**
     * Sample code: Get a community gallery image.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getACommunityGalleryImage(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCommunityGalleryImages().getWithResponse("myLocation", "publicGalleryName",
            "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
