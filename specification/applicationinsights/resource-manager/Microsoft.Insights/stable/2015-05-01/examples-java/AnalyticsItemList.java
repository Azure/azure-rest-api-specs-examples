import com.azure.resourcemanager.applicationinsights.models.ItemScopePath;

/** Samples for AnalyticsItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemList.json
     */
    /**
     * Sample code: AnalyticsItemList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void analyticsItemList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .analyticsItems()
            .listWithResponse(
                "my-resource-group",
                "my-component",
                ItemScopePath.ANALYTICS_ITEMS,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
