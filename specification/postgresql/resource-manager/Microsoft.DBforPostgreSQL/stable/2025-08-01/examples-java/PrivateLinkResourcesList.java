
/**
 * Samples for PrivateLinkResources ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * PrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for PostgreSQL.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getsPrivateLinkResourcesForPostgreSQL(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.privateLinkResources().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
