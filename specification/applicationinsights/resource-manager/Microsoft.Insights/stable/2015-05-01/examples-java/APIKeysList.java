/** Samples for ApiKeys List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysList.json
     */
    /**
     * Sample code: APIKeysList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void aPIKeysList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.apiKeys().list("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
