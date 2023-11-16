/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/ServerGetWithPrivateEndpoints.json
     */
    /**
     * Sample code: ServerGetWithPrivateEndpoints.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverGetWithPrivateEndpoints(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "pgtestsvc2", com.azure.core.util.Context.NONE);
    }
}
