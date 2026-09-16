
/**
 * Samples for EntityQueryTemplates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entityQueryTemplates/GetActivityEntityQueryTemplateById.json
     */
    /**
     * Sample code: Get an Activity entity query template.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnActivityEntityQueryTemplate(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entityQueryTemplates().getWithResponse("myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b",
            com.azure.core.util.Context.NONE);
    }
}
