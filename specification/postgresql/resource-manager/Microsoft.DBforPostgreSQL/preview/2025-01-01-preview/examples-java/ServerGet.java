
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ServerGet
     * .json
     */
    /**
     * Sample code: ServerGet.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverGet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "testpgflex", com.azure.core.util.Context.NONE);
    }
}
