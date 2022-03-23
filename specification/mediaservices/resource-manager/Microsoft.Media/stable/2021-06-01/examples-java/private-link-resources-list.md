Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/private-link-resources-list.json
     */
    /**
     * Sample code: Get list of all group IDs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getListOfAllGroupIDs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.privateLinkResources().listWithResponse("contoso", "contososports", Context.NONE);
    }
}
```
