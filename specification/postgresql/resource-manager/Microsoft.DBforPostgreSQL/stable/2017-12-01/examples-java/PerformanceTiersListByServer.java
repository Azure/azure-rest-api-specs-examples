import com.azure.core.util.Context;

/** Samples for ServerBasedPerformanceTier List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/PerformanceTiersListByServer.json
     */
    /**
     * Sample code: PerformanceTiersList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void performanceTiersList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverBasedPerformanceTiers().list("TestGroup", "testserver", Context.NONE);
    }
}
