
import com.azure.resourcemanager.storage.models.BlobRestoreParameters;
import com.azure.resourcemanager.storage.models.BlobRestoreRange;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for StorageAccounts RestoreBlobRanges.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobRangesRestore.json
     */
    /**
     * Sample code: BlobRangesRestore.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void blobRangesRestore(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().restoreBlobRanges("res9101", "sto4445",
            new BlobRestoreParameters().withTimeToRestore(OffsetDateTime.parse("2019-04-20T15:30:00.0000000Z"))
                .withBlobRanges(Arrays.asList(
                    new BlobRestoreRange().withStartRange("container/blobpath1").withEndRange("container/blobpath2"),
                    new BlobRestoreRange().withStartRange("container2/blobpath3").withEndRange(""))),
            com.azure.core.util.Context.NONE);
    }
}
