import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.DeletedShare;

/** Samples for FileShares Restore. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileSharesRestore.json
     */
    /**
     * Sample code: RestoreShares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restoreShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .restoreWithResponse(
                "res3376",
                "sto328",
                "share1249",
                new DeletedShare().withDeletedShareName("share1249").withDeletedShareVersion("1234567890"),
                Context.NONE);
    }
}
