
/**
 * Samples for StreamingLocators ListContentKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * locators-list-content-keys.json
     */
    /**
     * Sample code: List Content Keys.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listContentKeys(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingLocators().listContentKeysWithResponse("contosorg", "contosomedia", "secureStreamingLocator",
            com.azure.core.util.Context.NONE);
    }
}
