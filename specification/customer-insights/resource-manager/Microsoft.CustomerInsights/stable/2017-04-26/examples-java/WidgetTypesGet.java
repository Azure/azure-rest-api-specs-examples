/** Samples for WidgetTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/WidgetTypesGet.json
     */
    /**
     * Sample code: WidgetTypes_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void widgetTypesGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .widgetTypes()
            .getWithResponse("TestHubRG", "sdkTestHub", "ActivityGauge", com.azure.core.util.Context.NONE);
    }
}
