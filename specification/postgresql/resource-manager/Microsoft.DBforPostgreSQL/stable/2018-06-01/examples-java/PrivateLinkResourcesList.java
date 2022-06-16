import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for PostgreSQL.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getsPrivateLinkResourcesForPostgreSQL(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.privateLinkResources().listByServer("Default", "test-svr", Context.NONE);
    }
}
