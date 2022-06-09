```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.AssetTrack;
import com.azure.resourcemanager.mediaservices.models.TextTrack;

/** Samples for Tracks Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-update.json
     */
    /**
     * Sample code: Update a Track.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateATrack(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        AssetTrack resource =
            manager
                .tracks()
                .getWithResponse("contoso", "contosomedia", "ClimbingMountRainer", "text1", Context.NONE)
                .getValue();
        resource.update().withTrack(new TextTrack().withDisplayName("A new name")).apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
