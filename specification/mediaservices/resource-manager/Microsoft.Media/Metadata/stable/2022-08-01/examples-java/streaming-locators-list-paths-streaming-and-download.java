
/**
 * Samples for StreamingLocators ListPaths.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * locators-list-paths-streaming-and-download.json
     */
    /**
     * Sample code: List Paths which has streaming paths and download paths.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listPathsWhichHasStreamingPathsAndDownloadPaths(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingLocators().listPathsWithResponse("contosorg", "contosomedia", "clearStreamingLocator",
            com.azure.core.util.Context.NONE);
    }
}
