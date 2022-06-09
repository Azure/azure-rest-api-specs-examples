```java
import java.time.OffsetDateTime;

/** Samples for StreamingLocators Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-locators-create-secure.json
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
