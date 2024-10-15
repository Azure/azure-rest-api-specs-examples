
/**
 * Samples for ProductSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * settings/GetAllSettings.json
     */
    /**
     * Sample code: Get all settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productSettings().listWithResponse("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
