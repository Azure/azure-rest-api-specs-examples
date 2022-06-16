import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.EnabledProtocols;

/** Samples for FileShares Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/FileSharesPut_NFS.json
     */
    /**
     * Sample code: Create NFS Shares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNFSShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .createWithResponse(
                "res346",
                "sto666",
                "share1235",
                new FileShareInner().withEnabledProtocols(EnabledProtocols.NFS),
                null,
                Context.NONE);
    }
}
