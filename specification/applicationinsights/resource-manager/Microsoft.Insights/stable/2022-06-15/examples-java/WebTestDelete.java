/** Samples for WebTests Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestDelete.json
     */
    /**
     * Sample code: webTestDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestDelete(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .webTests()
            .deleteByResourceGroupWithResponse(
                "my-resource-group", "my-webtest-01-mywebservice", com.azure.core.util.Context.NONE);
    }
}
