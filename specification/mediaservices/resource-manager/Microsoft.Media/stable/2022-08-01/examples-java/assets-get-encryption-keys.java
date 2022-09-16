import com.azure.core.util.Context;

/** Samples for Assets GetEncryptionKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assets-get-encryption-keys.json
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
