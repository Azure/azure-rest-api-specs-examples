import com.azure.core.util.Context;

/** Samples for RecoverableServers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/RecoverableServersGet.json
     */
    /**
     * Sample code: ReplicasListByServer.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void replicasListByServer(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.recoverableServers().getWithResponse("testrg", "testsvc4", Context.NONE);
    }
}
