
/**
 * Samples for OfficeConsents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * officeConsents/GetOfficeConsentsById.json
     */
    /**
     * Sample code: Get an office consent.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnOfficeConsent(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.officeConsents().getWithResponse("myRg", "myWorkspace", "04e5fd05-ff86-4b97-b8d2-1c20933cb46c",
            com.azure.core.util.Context.NONE);
    }
}
