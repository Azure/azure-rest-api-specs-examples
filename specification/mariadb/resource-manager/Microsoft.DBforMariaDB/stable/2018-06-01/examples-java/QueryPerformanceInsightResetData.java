/** Samples for ResourceProvider ResetQueryPerformanceInsightData. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/QueryPerformanceInsightResetData.json
     */
    /**
     * Sample code: QueryPerformanceInsightResetData.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void queryPerformanceInsightResetData(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .resourceProviders()
            .resetQueryPerformanceInsightDataWithResponse(
                "testResourceGroupName", "testServerName", com.azure.core.util.Context.NONE);
    }
}
