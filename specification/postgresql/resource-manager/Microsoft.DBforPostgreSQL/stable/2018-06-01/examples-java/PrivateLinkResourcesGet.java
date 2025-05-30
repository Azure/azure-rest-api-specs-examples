
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/
     * PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for PostgreSQL.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getsAPrivateLinkResourceForPostgreSQL(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.privateLinkResources().getWithResponse("Default", "test-svr", "plr", com.azure.core.util.Context.NONE);
    }
}
