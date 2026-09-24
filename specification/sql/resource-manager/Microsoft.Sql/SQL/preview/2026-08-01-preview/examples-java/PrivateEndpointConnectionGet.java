
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("Default", "test-svr",
            "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
