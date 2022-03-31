Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Tracks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-delete.json
     */
    /**
     * Sample code: Delete a Track.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteATrack(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().delete("contoso", "contosomedia", "ClimbingMountRainer", "text2", Context.NONE);
    }
}
```
