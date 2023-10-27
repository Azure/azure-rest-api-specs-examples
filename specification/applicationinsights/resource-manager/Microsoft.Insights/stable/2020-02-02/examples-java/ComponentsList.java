/** Samples for Components List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsList.json
     */
    /**
     * Sample code: ComponentsList.json.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentsListJson(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.components().list(com.azure.core.util.Context.NONE);
    }
}
