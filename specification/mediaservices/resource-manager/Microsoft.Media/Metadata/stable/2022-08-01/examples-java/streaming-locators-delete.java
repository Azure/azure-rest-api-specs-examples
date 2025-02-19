
/**
 * Samples for StreamingLocators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * locators-delete.json
     */
    /**
     * Sample code: Delete a Streaming Locator.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAStreamingLocator(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingLocators().deleteWithResponse("contosorg", "contosomedia", "clearStreamingLocator",
            com.azure.core.util.Context.NONE);
    }
}
