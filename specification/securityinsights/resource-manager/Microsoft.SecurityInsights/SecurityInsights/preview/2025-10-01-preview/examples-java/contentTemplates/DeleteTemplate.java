
/**
 * Samples for ContentTemplate Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/contentTemplates/DeleteTemplate.json
     */
    /**
     * Sample code: Delete metadata.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.contentTemplates().deleteWithResponse("myRg", "myWorkspace", "8365ebfe-a381-45b7-ad08-7d818070e11f",
            com.azure.core.util.Context.NONE);
    }
}
