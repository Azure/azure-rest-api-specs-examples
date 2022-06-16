import com.azure.core.util.Context;

/** Samples for Servers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerDelete.json
     */
    /**
     * Sample code: ServerDelete.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverDelete(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().delete("TestGroup", "testserver", Context.NONE);
    }
}
