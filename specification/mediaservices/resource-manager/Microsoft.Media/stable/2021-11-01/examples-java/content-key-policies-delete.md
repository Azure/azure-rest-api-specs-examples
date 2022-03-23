Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContentKeyPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-delete.json
     */
    /**
     * Sample code: Delete a Key Policy.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAKeyPolicy(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .deleteWithResponse("contoso", "contosomedia", "PolicyWithPlayReadyOptionAndOpenRestriction", Context.NONE);
    }
}
```
