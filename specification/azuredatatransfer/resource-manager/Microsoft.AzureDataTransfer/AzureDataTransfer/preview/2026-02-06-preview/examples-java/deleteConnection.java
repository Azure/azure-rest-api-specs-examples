
/**
 * Samples for Connections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/deleteConnection.json
     */
    /**
     * Sample code: Deletes the connection resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        deletesTheConnectionResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().delete("testRG", "testConnection", com.azure.core.util.Context.NONE);
    }
}
