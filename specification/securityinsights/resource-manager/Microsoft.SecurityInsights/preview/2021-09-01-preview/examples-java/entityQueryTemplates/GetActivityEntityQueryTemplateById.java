import com.azure.core.util.Context;

/** Samples for EntityQueryTemplates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueryTemplates/GetActivityEntityQueryTemplateById.json
     */
    /**
     * Sample code: Get an Activity entity query template.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnActivityEntityQueryTemplate(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entityQueryTemplates()
            .getWithResponse("myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b", Context.NONE);
    }
}
