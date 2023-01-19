/** Samples for ServerBasedPerformanceTier List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/PerformanceTiersListByServer.json
     */
    /**
     * Sample code: PerformanceTiersList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void performanceTiersList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.serverBasedPerformanceTiers().list("testrg", "mariadbtestsvc1", com.azure.core.util.Context.NONE);
    }
}
