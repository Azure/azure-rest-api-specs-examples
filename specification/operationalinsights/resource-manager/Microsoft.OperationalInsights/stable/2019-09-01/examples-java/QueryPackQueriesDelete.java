import com.azure.core.util.Context;

/** Samples for Queries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesDelete.json
     */
    /**
     * Sample code: QueryDelete.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .queries()
            .deleteWithResponse(
                "my-resource-group", "my-querypack", "a449f8af-8e64-4b3a-9b16-5a7165ff98c4", Context.NONE);
    }
}
