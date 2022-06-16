import com.azure.core.util.Context;

/** Samples for ResourceProvider ResetQueryPerformanceInsightData. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/QueryPerformanceInsightResetData.json
     */
    /**
     * Sample code: QueryPerformanceInsightResetData.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void queryPerformanceInsightResetData(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .resourceProviders()
            .resetQueryPerformanceInsightDataWithResponse("testResourceGroupName", "testServerName", Context.NONE);
    }
}
