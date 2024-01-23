
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/
     * PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for PostgreSQL.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getsAPrivateLinkResourceForPostgreSQL(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.privateLinkResources().getWithResponse("Default", "test-svr", "plr", com.azure.core.util.Context.NONE);
    }
}
