import com.azure.core.util.Context;

/** Samples for Servers Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerRestart.json
     */
    /**
     * Sample code: ServerRestart.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverRestart(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().restart("TestGroup", "testserver", Context.NONE);
    }
}
