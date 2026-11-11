
/**
 * Samples for Connections ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Connections_ListAll.json
     */
    /**
     * Sample code: Get a list of connections.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAListOfConnections(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.connections().listAll("exampleRG", "exampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
