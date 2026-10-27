
/**
 * Samples for Connections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listPendingConnections.json
     */
    /**
     * Sample code: Lists a list of pending connections.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        listsAListOfPendingConnections(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().list("testRG", "testConnection", com.azure.core.util.Context.NONE);
    }
}
