
/**
 * Samples for VirtualEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * VirtualEndpointsGet.json
     */
    /**
     * Sample code: Get a virtual endpoint.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getAVirtualEndpoint(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.virtualEndpoints().getWithResponse("testrg", "pgtestsvc4", "pgVirtualEndpoint1",
            com.azure.core.util.Context.NONE);
    }
}
