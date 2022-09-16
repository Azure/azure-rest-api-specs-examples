import com.azure.core.util.Context;

/** Samples for StreamingLocators Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-locators-delete.json
     */
    /**
     * Sample code: Delete a Streaming Locator.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAStreamingLocator(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingLocators()
            .deleteWithResponse("contoso", "contosomedia", "clearStreamingLocator", Context.NONE);
    }
}
