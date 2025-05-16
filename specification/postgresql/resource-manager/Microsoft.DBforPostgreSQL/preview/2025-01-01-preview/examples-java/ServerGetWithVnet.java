
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * ServerGetWithVnet.json
     */
    /**
     * Sample code: ServerGetWithVnet.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverGetWithVnet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "testpgflex", com.azure.core.util.Context.NONE);
    }
}
