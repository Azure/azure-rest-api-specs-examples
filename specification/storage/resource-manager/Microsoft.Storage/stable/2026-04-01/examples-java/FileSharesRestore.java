
import com.azure.resourcemanager.storage.models.DeletedShare;

/**
 * Samples for FileShares Restore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileSharesRestore.json
     */
    /**
     * Sample code: RestoreShares.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void restoreShares(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().restoreWithResponse("res3376", "sto328", "share1249",
            new DeletedShare().withDeletedShareName("share1249").withDeletedShareVersion("1234567890"),
            com.azure.core.util.Context.NONE);
    }
}
