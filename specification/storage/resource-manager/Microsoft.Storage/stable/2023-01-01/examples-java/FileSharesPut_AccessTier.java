
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.ShareAccessTier;

/** Samples for FileShares Create. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesPut_AccessTier.json
     */
    /**
     * Sample code: PutShares with Access Tier.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putSharesWithAccessTier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().createWithResponse("res346", "sto666",
            "share1235", new FileShareInner().withAccessTier(ShareAccessTier.HOT), null,
            com.azure.core.util.Context.NONE);
    }
}
