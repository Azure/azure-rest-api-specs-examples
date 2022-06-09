```java
import com.azure.core.util.Context;

/** Samples for Mediaservices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-delete.json
     */
    /**
     * Sample code: Delete a Media Services account.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAMediaServicesAccount(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.mediaservices().deleteWithResponse("contoso", "contososports", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
