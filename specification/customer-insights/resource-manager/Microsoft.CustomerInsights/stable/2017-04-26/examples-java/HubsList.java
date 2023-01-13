/** Samples for Hubs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsList.json
     */
    /**
     * Sample code: Hubs_List.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void hubsList(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.hubs().list(com.azure.core.util.Context.NONE);
    }
}
