
import com.azure.resourcemanager.datatransfer.models.ConnectionsPatch;

/**
 * Samples for Connections Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/patchConnection.json
     */
    /**
     * Sample code: Updates the connection resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        updatesTheConnectionResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().update("testRG", "testConnection", new ConnectionsPatch(),
            com.azure.core.util.Context.NONE);
    }
}
