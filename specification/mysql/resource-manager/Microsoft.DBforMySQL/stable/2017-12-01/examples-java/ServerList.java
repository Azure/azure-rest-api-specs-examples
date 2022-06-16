import com.azure.core.util.Context;

/** Samples for Servers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerList.json
     */
    /**
     * Sample code: ServerList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().list(Context.NONE);
    }
}
