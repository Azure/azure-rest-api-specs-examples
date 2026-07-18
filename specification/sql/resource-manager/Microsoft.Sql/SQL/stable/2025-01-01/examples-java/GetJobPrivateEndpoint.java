
/**
 * Samples for JobPrivateEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobPrivateEndpoint.json
     */
    /**
     * Sample code: Get a private endpoint.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAPrivateEndpoint(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobPrivateEndpoints().getWithResponse("group1", "server1", "agent1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
