
/**
 * Samples for ProductSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * settings/DeleteEyesOnSetting.json
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
