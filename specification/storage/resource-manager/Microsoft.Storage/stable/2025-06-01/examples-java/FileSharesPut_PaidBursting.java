
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.FileSharePropertiesFileSharePaidBursting;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/FileSharesPut_PaidBursting.
     * json
     */
    /**
     * Sample code: PutShares with Paid Bursting.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putSharesWithPaidBursting(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().createWithResponse("res346", "sto666",
            "share1235",
            new FileShareInner()
                .withFileSharePaidBursting(new FileSharePropertiesFileSharePaidBursting().withPaidBurstingEnabled(true)
                    .withPaidBurstingMaxIops(102400).withPaidBurstingMaxBandwidthMibps(10340)),
            null, com.azure.core.util.Context.NONE);
    }
}
