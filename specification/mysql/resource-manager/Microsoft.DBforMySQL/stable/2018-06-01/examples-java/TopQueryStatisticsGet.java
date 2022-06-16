import com.azure.core.util.Context;

/** Samples for TopQueryStatistics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/TopQueryStatisticsGet.json
     */
    /**
     * Sample code: TopQueryStatisticsGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void topQueryStatisticsGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .topQueryStatistics()
            .getWithResponse(
                "testResourceGroupName",
                "testServerName",
                "66-636923268000000000-636923277000000000-avg-duration",
                Context.NONE);
    }
}
