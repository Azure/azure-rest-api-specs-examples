
/**
 * Samples for StreamingLocators List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * locators-list.json
     */
    /**
     * Sample code: Lists Streaming Locators.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsStreamingLocators(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingLocators().list("contosorg", "contosomedia", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
