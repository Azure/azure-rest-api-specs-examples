
/**
 * Samples for Images ListDownloadUri.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/Images_ListDownloadUri_MaximumSet_Gen.json
     */
    /**
     * Sample code: Images_ListDownloadUri.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void
        imagesListDownloadUri(com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.images().listDownloadUriWithResponse("rgdisconnectedOperations", "g_-5-160", "1Q6lGV4V65j-1",
            com.azure.core.util.Context.NONE);
    }
}
