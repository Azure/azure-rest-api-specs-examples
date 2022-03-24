Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Assets ListStreamingLocators. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-streaming-locators.json
     */
    /**
     * Sample code: List Asset SAS URLs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAssetSASURLs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .listStreamingLocatorsWithResponse("contoso", "contosomedia", "ClimbingMountSaintHelens", Context.NONE);
    }
}
```
