Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Tracks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-get-by-name.json
     */
    /**
     * Sample code: Get a Track by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getATrackByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().getWithResponse("contoso", "contosomedia", "ClimbingMountRainer", "text1", Context.NONE);
    }
}
```
