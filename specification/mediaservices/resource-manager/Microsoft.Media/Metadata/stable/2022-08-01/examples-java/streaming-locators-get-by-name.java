/** Samples for StreamingLocators Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-get-by-name.json
     */
    /**
     * Sample code: Get a Streaming Locator by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAStreamingLocatorByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingLocators()
            .getWithResponse("contoso", "contosomedia", "clearStreamingLocator", com.azure.core.util.Context.NONE);
    }
}
