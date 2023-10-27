/** Samples for Annotations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsGet.json
     */
    /**
     * Sample code: AnnotationsGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void annotationsGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .annotations()
            .getWithResponse(
                "my-resource-group",
                "my-component",
                "444e2c08-274a-4bbb-a89e-d77bb720f44a",
                com.azure.core.util.Context.NONE);
    }
}
