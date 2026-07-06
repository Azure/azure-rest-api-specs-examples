
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.FileSharePropertiesFileSharePaidBursting;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileSharesPut_PaidBursting.json
     */
    /**
     * Sample code: PutShares with Paid Bursting.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putSharesWithPaidBursting(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().createWithResponse("res346", "sto666", "share1235",
            new FileShareInner()
                .withFileSharePaidBursting(new FileSharePropertiesFileSharePaidBursting().withPaidBurstingEnabled(true)
                    .withPaidBurstingMaxIops(102400).withPaidBurstingMaxBandwidthMibps(10340)),
            null, com.azure.core.util.Context.NONE);
    }
}
