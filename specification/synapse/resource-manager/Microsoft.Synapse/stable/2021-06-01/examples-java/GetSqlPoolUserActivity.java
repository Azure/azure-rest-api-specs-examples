import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.DataWarehouseUserActivityName;

/** Samples for SqlPoolDataWarehouseUserActivities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolUserActivity.json
     */
    /**
     * Sample code: Get a SQL Analytics pool user activity.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getASQLAnalyticsPoolUserActivity(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolDataWarehouseUserActivities()
            .getWithResponse(
                "Default-SQL-SouthEastAsia", "testsvr", "testdb", DataWarehouseUserActivityName.CURRENT, Context.NONE);
    }
}
