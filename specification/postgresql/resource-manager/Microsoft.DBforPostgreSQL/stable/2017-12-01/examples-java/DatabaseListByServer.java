import com.azure.core.util.Context;

/** Samples for Databases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseListByServer.json
     */
    /**
     * Sample code: DatabaseList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void databaseList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.databases().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
