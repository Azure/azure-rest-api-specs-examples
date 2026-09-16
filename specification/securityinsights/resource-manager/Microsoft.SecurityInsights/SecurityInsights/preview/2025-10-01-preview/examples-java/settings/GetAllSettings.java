
/**
 * Samples for ProductSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/settings/GetAllSettings.json
     */
    /**
     * Sample code: Get all settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productSettings().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
