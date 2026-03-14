
/**
 * Samples for Schedulers GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PrivateEndpointConnections_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        privateEndpointConnectionsGetMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().getPrivateEndpointConnectionWithResponse("rgdurabletask", "testscheduler",
            "spzckqrbhfnabu", com.azure.core.util.Context.NONE);
    }
}
