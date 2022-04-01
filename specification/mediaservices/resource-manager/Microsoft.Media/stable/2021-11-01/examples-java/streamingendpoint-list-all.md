Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StreamingEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-list-all.json
     */
    /**
     * Sample code: List all streaming endpoints.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllStreamingEndpoints(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().list("mediaresources", "slitestmedia10", Context.NONE);
    }
}
```
