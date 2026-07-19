
/**
 * Samples for ManagedInstancePrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstancePrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstancePrivateEndpointConnections().getWithResponse("Default", "test-cl",
            "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
