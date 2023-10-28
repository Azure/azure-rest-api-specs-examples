/** Samples for WebTests ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestListByResourceGroup.json
     */
    /**
     * Sample code: webTestListByResourceGroup.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestListByResourceGroup(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().listByResourceGroup("my-resource-group", com.azure.core.util.Context.NONE);
    }
}
