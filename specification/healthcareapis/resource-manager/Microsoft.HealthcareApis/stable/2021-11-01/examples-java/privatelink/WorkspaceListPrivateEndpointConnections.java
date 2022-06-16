import com.azure.core.util.Context;

/** Samples for WorkspacePrivateEndpointConnections ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceListPrivateEndpointConnections.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnection_List.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateEndpointConnectionList(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspacePrivateEndpointConnections().listByWorkspace("testRG", "workspace1", Context.NONE);
    }
}
