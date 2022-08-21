/** Samples for PrivateEndpointConnections Approve. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/PrivateEndpointConnections_Approve.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Approve.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void privateEndpointConnectionsApprove(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager
            .privateEndpointConnections()
            .define("myConnection")
            .withExistingGrafana("myResourceGroup", "myWorkspace")
            .create();
    }
}
