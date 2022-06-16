import com.azure.core.util.Context;

/** Samples for WaitStatistics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/WaitStatisticsGet.json
     */
    /**
     * Sample code: WaitStatisticsGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void waitStatisticsGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .waitStatistics()
            .getWithResponse(
                "testResourceGroupName",
                "testServerName",
                "636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0",
                Context.NONE);
    }
}
