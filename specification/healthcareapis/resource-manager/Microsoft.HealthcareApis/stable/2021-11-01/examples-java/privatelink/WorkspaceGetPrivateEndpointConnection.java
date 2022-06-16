import com.azure.core.util.Context;

/** Samples for WorkspacePrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnection_GetConnection.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateEndpointConnectionGetConnection(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .workspacePrivateEndpointConnections()
            .getWithResponse("testRG", "workspace1", "myConnection", Context.NONE);
    }
}
