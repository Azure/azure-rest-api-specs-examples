
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;
import com.azure.resourcemanager.storage.models.EnabledProtocols;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileSharesPut_NFS.json
     */
    /**
     * Sample code: Create NFS Shares.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createNFSShares(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().createWithResponse("res346", "sto666", "share1235",
            new FileShareInner().withEnabledProtocols(EnabledProtocols.NFS), null, com.azure.core.util.Context.NONE);
    }
}
