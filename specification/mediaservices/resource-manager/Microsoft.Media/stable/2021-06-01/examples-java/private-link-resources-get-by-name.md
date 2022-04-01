Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/private-link-resources-get-by-name.json
     */
    /**
     * Sample code: Get details of a group ID.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getDetailsOfAGroupID(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.privateLinkResources().getWithResponse("contoso", "contososports", "keydelivery", Context.NONE);
    }
}
```
