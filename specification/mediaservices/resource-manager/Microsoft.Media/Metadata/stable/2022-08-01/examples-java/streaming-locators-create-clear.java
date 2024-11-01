
/**
 * Samples for StreamingLocators Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-
     * locators-create-clear.json
     */
    /**
     * Sample code: Creates a Streaming Locator with clear streaming.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingLocatorWithClearStreaming(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingLocators().define("UserCreatedClearStreamingLocator")
            .withExistingMediaService("contosorg", "contosomedia").withAssetName("ClimbingMountRainier")
            .withStreamingPolicyName("clearStreamingPolicy").create();
    }
}
