
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesPut_ProvisionedV2.
     * json
     */
    /**
     * Sample code: PutSharesProvisionedV2.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putSharesProvisionedV2(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().createWithResponse("res346", "sto666",
            "share1235",
            new FileShareInner().withShareQuota(100).withProvisionedIops(5000).withProvisionedBandwidthMibps(200), null,
            com.azure.core.util.Context.NONE);
    }
}
