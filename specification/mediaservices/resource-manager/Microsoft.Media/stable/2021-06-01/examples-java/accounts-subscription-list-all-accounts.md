Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Mediaservices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-subscription-list-all-accounts.json
     */
    /**
     * Sample code: List all Media Services accounts.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllMediaServicesAccounts(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.mediaservices().list(Context.NONE);
    }
}
```
