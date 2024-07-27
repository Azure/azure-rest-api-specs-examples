
import com.azure.resourcemanager.storage.fluent.models.BlobServicePropertiesInner;
import com.azure.resourcemanager.storage.models.DeleteRetentionPolicy;

/**
 * Samples for BlobServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * BlobServicesPutAllowPermanentDelete.json
     */
    /**
     * Sample code: BlobServicesPutAllowPermanentDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void blobServicesPutAllowPermanentDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobServices().setServicePropertiesWithResponse("res4410",
            "sto8607",
            new BlobServicePropertiesInner()
                .withDeleteRetentionPolicy(
                    new DeleteRetentionPolicy().withEnabled(true).withDays(300).withAllowPermanentDelete(true))
                .withIsVersioningEnabled(true),
            com.azure.core.util.Context.NONE);
    }
}
