
/**
 * Samples for CommunityGalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/communityGalleryExamples/CommunityGalleryImageVersion_Get.json
     */
    /**
     * Sample code: Get a community gallery image version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getACommunityGalleryImageVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCommunityGalleryImageVersions().getWithResponse("myLocation", "publicGalleryName",
            "myGalleryImageName", "myGalleryImageVersionName", com.azure.core.util.Context.NONE);
    }
}
