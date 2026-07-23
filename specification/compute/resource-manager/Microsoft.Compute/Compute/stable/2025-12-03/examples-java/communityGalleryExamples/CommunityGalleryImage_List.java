
/**
 * Samples for CommunityGalleryImages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/communityGalleryExamples/CommunityGalleryImage_List.json
     */
    /**
     * Sample code: List community gallery images.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listCommunityGalleryImages(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getCommunityGalleryImages().list("myLocation", "publicGalleryName",
            com.azure.core.util.Context.NONE);
    }
}
