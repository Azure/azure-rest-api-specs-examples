/** Samples for ApiKeys Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysGet.json
     */
    /**
     * Sample code: APIKeysGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void aPIKeysGet(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .apiKeys()
            .getWithResponse(
                "my-resource-group",
                "my-component",
                "bb820f1b-3110-4a8b-ba2c-8c1129d7eb6a",
                com.azure.core.util.Context.NONE);
    }
}
