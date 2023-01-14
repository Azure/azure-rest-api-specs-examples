/** Samples for WidgetTypes ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/WidgetTypesListByHub.json
     */
    /**
     * Sample code: WidgetTypes_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void widgetTypesListByHub(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.widgetTypes().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
