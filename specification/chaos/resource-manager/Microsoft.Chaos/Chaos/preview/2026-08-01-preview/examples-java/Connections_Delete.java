
/**
 * Samples for Connections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/Connections_Delete.json
     */
    /**
     * Sample code: Delete a connection in a workspace.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAConnectionInAWorkspace(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.connections().deleteWithResponse("exampleRG", "exampleWorkspace", "aksClusterConnection",
            com.azure.core.util.Context.NONE);
    }
}
