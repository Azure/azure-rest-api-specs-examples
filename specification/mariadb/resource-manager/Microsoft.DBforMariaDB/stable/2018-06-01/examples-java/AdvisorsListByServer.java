/** Samples for Advisors ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/AdvisorsListByServer.json
     */
    /**
     * Sample code: AdvisorsListByServer.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void advisorsListByServer(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.advisors().listByServer("testResourceGroupName", "testServerName", com.azure.core.util.Context.NONE);
    }
}
