/** Samples for LocationBasedPerformanceTier List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/PerformanceTiersListByLocation.json
     */
    /**
     * Sample code: PerformanceTiersList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void performanceTiersList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.locationBasedPerformanceTiers().list("WestUS", com.azure.core.util.Context.NONE);
    }
}
