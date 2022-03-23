Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Assets GetEncryptionKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-get-encryption-keys.json
     */
    /**
     * Sample code: Get Asset Storage Encryption Keys.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAssetStorageEncryptionKeys(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .getEncryptionKeyWithResponse("contoso", "contosomedia", "ClimbingMountSaintHelens", Context.NONE);
    }
}
```
