
/**
 * Samples for ProductSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * settings/GetEyesOnSetting.json
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
