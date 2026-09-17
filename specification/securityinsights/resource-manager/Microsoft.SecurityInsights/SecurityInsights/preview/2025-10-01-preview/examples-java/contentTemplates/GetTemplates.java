
/**
 * Samples for ContentTemplatesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentTemplates/GetTemplates.json
     */
    /**
     * Sample code: Get all installed templates.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllInstalledTemplates(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.contentTemplatesOperations().list("myRg", "myWorkspace", null, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
