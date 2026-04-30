
/**
 * Samples for SharedGalleries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/sharedGalleryExamples/SharedGallery_List.json
     */
    /**
     * Sample code: List shared galleries.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listSharedGalleries(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleries().list("myLocation", null, com.azure.core.util.Context.NONE);
    }
}
