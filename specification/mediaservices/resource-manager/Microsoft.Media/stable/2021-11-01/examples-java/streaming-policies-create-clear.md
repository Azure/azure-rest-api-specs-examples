Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.NoEncryption;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-create-clear.json
     */
    /**
     * Sample code: Creates a Streaming Policy with clear streaming.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithClearStreaming(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingPolicies()
            .define("UserCreatedClearStreamingPolicy")
            .withExistingMediaService("contoso", "contosomedia")
            .withNoEncryption(
                new NoEncryption()
                    .withEnabledProtocols(
                        new EnabledProtocols()
                            .withDownload(true)
                            .withDash(true)
                            .withHls(true)
                            .withSmoothStreaming(true)))
            .create();
    }
}
```
