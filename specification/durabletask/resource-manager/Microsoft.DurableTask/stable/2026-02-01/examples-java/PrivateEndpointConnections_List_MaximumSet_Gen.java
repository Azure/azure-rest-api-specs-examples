
/**
 * Samples for Schedulers ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PrivateEndpointConnections_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        privateEndpointConnectionsListMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().listPrivateEndpointConnections("rgdurabletask", "testscheduler",
            com.azure.core.util.Context.NONE);
    }
}
