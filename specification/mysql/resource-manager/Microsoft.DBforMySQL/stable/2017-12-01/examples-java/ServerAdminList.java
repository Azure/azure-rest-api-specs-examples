import com.azure.core.util.Context;

/** Samples for ServerAdministrators List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminList.json
     */
    /**
     * Sample code: get a list of server administrators.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAListOfServerAdministrators(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.serverAdministrators().list("testrg", "mysqltestsvc4", Context.NONE);
    }
}
