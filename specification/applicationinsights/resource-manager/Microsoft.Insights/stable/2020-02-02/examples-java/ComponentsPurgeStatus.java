/** Samples for Components GetPurgeStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsPurgeStatus.json
     */
    /**
     * Sample code: ComponentPurge.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentPurge(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .components()
            .getPurgeStatusWithResponse(
                "OIAutoRest5123",
                "aztest5048",
                "purge-970318e7-b859-4edb-8903-83b1b54d0b74",
                com.azure.core.util.Context.NONE);
    }
}
