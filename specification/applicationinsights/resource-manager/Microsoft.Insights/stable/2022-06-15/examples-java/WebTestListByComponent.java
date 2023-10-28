/** Samples for WebTests ListByComponent. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestListByComponent.json
     */
    /**
     * Sample code: webTestListByComponent.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestListByComponent(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().listByComponent("my-component", "my-resource-group", com.azure.core.util.Context.NONE);
    }
}
