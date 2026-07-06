
import com.azure.resourcemanager.storage.fluent.models.BlobServicePropertiesInner;
import com.azure.resourcemanager.storage.models.DeleteRetentionPolicy;

/**
 * Samples for BlobServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobServicesPutAllowPermanentDelete.json
     */
    /**
     * Sample code: BlobServicesPutAllowPermanentDelete.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void blobServicesPutAllowPermanentDelete(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobServices().setServicePropertiesWithResponse("res4410", "sto8607",
            new BlobServicePropertiesInner()
                .withDeleteRetentionPolicy(
                    new DeleteRetentionPolicy().withEnabled(true).withDays(300).withAllowPermanentDelete(true))
                .withIsVersioningEnabled(true),
            com.azure.core.util.Context.NONE);
    }
}
