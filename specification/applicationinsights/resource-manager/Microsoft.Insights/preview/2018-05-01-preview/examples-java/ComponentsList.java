import com.azure.core.util.Context;

/** Samples for Components List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsList.json
     */
    /**
     * Sample code: ComponentsList.json.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentsListJson(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.components().list(Context.NONE);
    }
}
