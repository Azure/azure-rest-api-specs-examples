import com.azure.core.util.Context;

/** Samples for Servers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerDelete.json
     */
    /**
     * Sample code: ServerDelete.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverDelete(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.servers().delete("TestGroup", "testserver", Context.NONE);
    }
}
