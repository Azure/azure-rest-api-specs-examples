
/**
 * Samples for VirtualEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/VirtualEndpointsGet.json
     */
    /**
     * Sample code: Get information about a pair of virtual endpoints.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAPairOfVirtualEndpoints(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.virtualEndpoints().getWithResponse("exampleresourcegroup", "exampleserver", "examplebasename",
            com.azure.core.util.Context.NONE);
    }
}
