import com.azure.core.util.Context;

/** Samples for ServerAdministrators Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerAdminGet.json
     */
    /**
     * Sample code: ServerAdministratorGet.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverAdministratorGet(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverAdministrators().getWithResponse("testrg", "pgtestsvc4", Context.NONE);
    }
}
