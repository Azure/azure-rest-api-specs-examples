/** Samples for WebTests GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestGet.json
     */
    /**
     * Sample code: webTestGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestGet(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .webTests()
            .getByResourceGroupWithResponse(
                "my-resource-group", "my-webtest-01-mywebservice", com.azure.core.util.Context.NONE);
    }
}
