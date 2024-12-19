
/**
 * Samples for WorkspacePrivateEndpointConnections ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/privatelink/
     * WorkspaceListPrivateEndpointConnections.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnection_List.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        workspacePrivateEndpointConnectionList(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspacePrivateEndpointConnections().listByWorkspace("testRG", "workspace1",
            com.azure.core.util.Context.NONE);
    }
}
