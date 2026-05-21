
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;

/**
 * Samples for FileShares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesPut.json
     */
    /**
     * Sample code: PutShares.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putShares(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().createWithResponse("res3376", "sto328", "share6185",
            new FileShareInner(), null, com.azure.core.util.Context.NONE);
    }
}
