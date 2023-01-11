/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ServerGet.json
     */
    /**
     * Sample code: ServerGet.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverGet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "pgtestsvc1", com.azure.core.util.Context.NONE);
    }
}
