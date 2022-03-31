Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AssetFilters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-list-all.json
     */
    /**
     * Sample code: List all Asset Filters.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllAssetFilters(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assetFilters().list("contoso", "contosomedia", "ClimbingMountRainer", Context.NONE);
    }
}
```
