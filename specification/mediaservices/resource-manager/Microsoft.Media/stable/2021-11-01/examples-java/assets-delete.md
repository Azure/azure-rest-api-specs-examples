Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Assets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-delete.json
     */
    /**
     * Sample code: Delete an Asset.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAnAsset(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assets().deleteWithResponse("contoso", "contosomedia", "ClimbingMountAdams", Context.NONE);
    }
}
```
