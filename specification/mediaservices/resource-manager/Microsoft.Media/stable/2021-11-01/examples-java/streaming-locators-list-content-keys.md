```java
import com.azure.core.util.Context;

/** Samples for StreamingLocators ListContentKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-locators-list-content-keys.json
     */
    /**
     * Sample code: List Content Keys.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listContentKeys(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingLocators()
            .listContentKeysWithResponse("contoso", "contosomedia", "secureStreamingLocator", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
