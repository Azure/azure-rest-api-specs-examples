import com.azure.core.util.Context;

/** Samples for ServerBasedPerformanceTier List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/PerformanceTiersListByServer.json
     */
    /**
     * Sample code: PerformanceTiersList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void performanceTiersList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.serverBasedPerformanceTiers().list("testrg", "mysqltestsvc1", Context.NONE);
    }
}
