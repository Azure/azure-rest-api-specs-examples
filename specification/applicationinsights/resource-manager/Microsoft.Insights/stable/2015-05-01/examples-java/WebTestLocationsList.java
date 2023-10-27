/** Samples for WebTestLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestLocationsList.json
     */
    /**
     * Sample code: WebTestLocationsList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestLocationsList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTestLocations().list("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
