import com.azure.core.util.Context;

/** Samples for WebTests ListByComponent. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestListByComponent.json
     */
    /**
     * Sample code: webTestListByComponent.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestListByComponent(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().listByComponent("my-component", "my-resource-group", Context.NONE);
    }
}
