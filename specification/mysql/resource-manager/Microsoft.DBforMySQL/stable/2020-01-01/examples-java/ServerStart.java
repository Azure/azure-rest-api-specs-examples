import com.azure.core.util.Context;

/** Samples for Servers Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerStart.json
     */
    /**
     * Sample code: ServerStart.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverStart(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().start("TestGroup", "testserver", Context.NONE);
    }
}
