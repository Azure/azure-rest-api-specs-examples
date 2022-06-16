import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.ItemScopePath;

/** Samples for AnalyticsItems Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemGet.json
     */
    /**
     * Sample code: AnalyticsItemGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void analyticsItemGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .analyticsItems()
            .getWithResponse(
                "my-resource-group",
                "my-component",
                ItemScopePath.ANALYTICS_ITEMS,
                "3466c160-4a10-4df8-afdf-0007f3f6dee5",
                null,
                Context.NONE);
    }
}
