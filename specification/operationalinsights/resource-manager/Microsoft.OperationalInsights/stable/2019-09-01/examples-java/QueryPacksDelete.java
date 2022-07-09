import com.azure.core.util.Context;

/** Samples for QueryPacks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksDelete.json
     */
    /**
     * Sample code: QueryPacksDelete.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPacksDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().deleteWithResponse("my-resource-group", "my-querypack", Context.NONE);
    }
}
