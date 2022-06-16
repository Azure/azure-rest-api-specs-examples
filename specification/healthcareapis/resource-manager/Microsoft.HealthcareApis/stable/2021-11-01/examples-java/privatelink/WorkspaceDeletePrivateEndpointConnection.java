import com.azure.core.util.Context;

/** Samples for WorkspacePrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateEndpointConnectionsDelete(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspacePrivateEndpointConnections().delete("testRG", "workspace1", "myConnection", Context.NONE);
    }
}
