import com.azure.core.util.Context;

/** Samples for Servers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerList.json
     */
    /**
     * Sample code: ServerList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.servers().list(Context.NONE);
    }
}
