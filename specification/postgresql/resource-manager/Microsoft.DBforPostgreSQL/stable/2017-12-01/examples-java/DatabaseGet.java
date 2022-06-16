import com.azure.core.util.Context;

/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseGet.json
     */
    /**
     * Sample code: DatabaseGet.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void databaseGet(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.databases().getWithResponse("TestGroup", "testserver", "db1", Context.NONE);
    }
}
