
/**
 * Samples for Flows ListByConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listFlowsByConnection.json
     */
    /**
     * Sample code: Gets flows in a connection.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void getsFlowsInAConnection(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().listByConnection("testRG", "testConnection", com.azure.core.util.Context.NONE);
    }
}
