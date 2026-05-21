
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesPut_ProvisionedV2.json
     */
    /**
     * Sample code: PutSharesProvisionedV2.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putSharesProvisionedV2(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().createWithResponse("res346", "sto666", "share1235",
            new FileShareInner().withShareQuota(100).withProvisionedIops(5000).withProvisionedBandwidthMibps(200), null,
            com.azure.core.util.Context.NONE);
    }
}
