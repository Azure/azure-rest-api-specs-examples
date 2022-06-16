import com.azure.core.util.Context;

/** Samples for RecoverableServers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/RecoverableServersGet.json
     */
    /**
     * Sample code: ReplicasListByServer.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void replicasListByServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.recoverableServers().getWithResponse("testrg", "pgtestsvc4", Context.NONE);
    }
}
