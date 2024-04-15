
/**
 * Samples for WorkspacePrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/privatelink/
     * WorkspaceGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnection_GetConnection.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateEndpointConnectionGetConnection(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspacePrivateEndpointConnections().getWithResponse("testRG", "workspace1", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
