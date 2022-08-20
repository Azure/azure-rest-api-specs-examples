import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateEndpointConnections().delete("myResourceGroup", "myWorkspace", "myConnection", Context.NONE);
    }
}
