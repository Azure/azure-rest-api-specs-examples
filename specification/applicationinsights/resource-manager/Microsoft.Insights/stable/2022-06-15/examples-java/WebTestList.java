/** Samples for WebTests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestList.json
     */
    /**
     * Sample code: webTestList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().list(com.azure.core.util.Context.NONE);
    }
}
