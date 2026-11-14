
/**
 * Samples for Location ListCachedImages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/CachedImagesList.json
     */
    /**
     * Sample code: CachedImages.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void cachedImages(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getLocations().listCachedImages("westcentralus", com.azure.core.util.Context.NONE);
    }
}
