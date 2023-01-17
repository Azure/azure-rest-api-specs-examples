import java.time.OffsetDateTime;

/** Samples for StreamingLocators Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure.json
     */
    /**
     * Sample code: Creates a Streaming Locator with secure streaming.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingLocatorWithSecureStreaming(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingLocators()
            .define("UserCreatedSecureStreamingLocator")
            .withExistingMediaService("contoso", "contosomedia")
            .withAssetName("ClimbingMountRainier")
            .withStartTime(OffsetDateTime.parse("2018-03-01T00:00:00Z"))
            .withEndTime(OffsetDateTime.parse("2028-12-31T23:59:59.9999999Z"))
            .withStreamingPolicyName("secureStreamingPolicy")
            .create();
    }
}
