
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.ShareAccessTier;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileSharesPut_AccessTier.json
     */
    /**
     * Sample code: PutShares with Access Tier.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putSharesWithAccessTier(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().createWithResponse("res346", "sto666", "share1235",
            new FileShareInner().withAccessTier(ShareAccessTier.HOT), null, com.azure.core.util.Context.NONE);
    }
}
