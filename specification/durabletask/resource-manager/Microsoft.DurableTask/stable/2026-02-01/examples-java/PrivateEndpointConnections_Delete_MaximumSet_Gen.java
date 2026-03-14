
/**
 * Samples for Schedulers DeletePrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PrivateEndpointConnections_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        privateEndpointConnectionsDeleteMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().deletePrivateEndpointConnection("rgdurabletask", "testscheduler", "spzckqrbhfnabu",
            com.azure.core.util.Context.NONE);
    }
}
