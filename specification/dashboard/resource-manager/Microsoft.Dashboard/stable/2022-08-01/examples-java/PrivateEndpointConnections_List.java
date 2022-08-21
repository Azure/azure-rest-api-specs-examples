import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/PrivateEndpointConnections_List.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateEndpointConnections().list("myResourceGroup", "myWorkspace", Context.NONE);
    }
}
