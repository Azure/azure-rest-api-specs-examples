/** Samples for RecoverableServers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecoverableServersGet.json
     */
    /**
     * Sample code: ReplicasListByServer.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void replicasListByServer(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.recoverableServers().getWithResponse("testrg", "testsvc4", com.azure.core.util.Context.NONE);
    }
}
