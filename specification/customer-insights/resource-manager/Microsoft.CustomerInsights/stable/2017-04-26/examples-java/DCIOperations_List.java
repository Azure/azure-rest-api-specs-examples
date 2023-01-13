/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/DCIOperations_List.json
     */
    /**
     * Sample code: DCIOperations_List.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void dCIOperationsList(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
