
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * ServerGetWithPrivateEndpoints.json
     */
    /**
     * Sample code: ServerGetWithPrivateEndpoints.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        serverGetWithPrivateEndpoints(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "pgtestsvc2", com.azure.core.util.Context.NONE);
    }
}
