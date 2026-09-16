
/**
 * Samples for ContentTemplate Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentTemplates/GetTemplateById.json
     */
    /**
     * Sample code: Get a template.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getATemplate(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.contentTemplates().getWithResponse("myRg", "myWorkspace", "8365ebfe-a381-45b7-ad08-7d818070e11f",
            com.azure.core.util.Context.NONE);
    }
}
