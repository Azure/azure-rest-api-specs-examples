import com.azure.core.util.Context;

/** Samples for ServerAdministrators Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminGet.json
     */
    /**
     * Sample code: ServerAdministratorGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverAdministratorGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.serverAdministrators().getWithResponse("testrg", "mysqltestsvc4", Context.NONE);
    }
}
