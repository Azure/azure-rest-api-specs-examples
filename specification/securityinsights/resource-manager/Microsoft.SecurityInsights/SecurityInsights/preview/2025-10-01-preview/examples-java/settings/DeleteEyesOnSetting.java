
/**
 * Samples for ProductSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/settings/DeleteEyesOnSetting.json
     */
    /**
     * Sample code: Delete EyesOn settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteEyesOnSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productSettings().deleteWithResponse("myRg", "myWorkspace", "EyesOn", com.azure.core.util.Context.NONE);
    }
}
