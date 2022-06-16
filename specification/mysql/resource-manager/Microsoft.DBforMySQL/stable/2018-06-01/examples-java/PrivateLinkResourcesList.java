import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for MySQL.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getsPrivateLinkResourcesForMySQL(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.privateLinkResources().listByServer("Default", "test-svr", Context.NONE);
    }
}
