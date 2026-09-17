
/**
 * Samples for ProductSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/settings/GetEyesOnSetting.json
     */
    /**
     * Sample code: Get EyesOn settings.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getEyesOnSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productSettings().getWithResponse("myRg", "myWorkspace", "EyesOn", com.azure.core.util.Context.NONE);
    }
}
