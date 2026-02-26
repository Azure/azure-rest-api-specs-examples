
/**
 * Samples for VirtualEndpoints ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/VirtualEndpointsListByServer.json
     */
    /**
     * Sample code: List pair of virtual endpoints associated to a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listPairOfVirtualEndpointsAssociatedToAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.virtualEndpoints().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
