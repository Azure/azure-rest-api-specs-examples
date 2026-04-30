
/**
 * Samples for Galleries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/Gallery_ListBySubscription.json
     */
    /**
     * Sample code: List galleries in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listGalleriesInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().list(com.azure.core.util.Context.NONE);
    }
}
