import com.azure.core.util.Context;

/** Samples for WebTests Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestDelete.json
     */
    /**
     * Sample code: webTestDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestDelete(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.webTests().deleteWithResponse("my-resource-group", "my-webtest-01-mywebservice", Context.NONE);
    }
}
