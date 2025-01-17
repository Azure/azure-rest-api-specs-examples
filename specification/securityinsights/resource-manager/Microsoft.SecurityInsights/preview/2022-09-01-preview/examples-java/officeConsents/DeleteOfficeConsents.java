
/**
 * Samples for OfficeConsents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * officeConsents/DeleteOfficeConsents.json
     */
    /**
     * Sample code: Delete an office consent.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAnOfficeConsent(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.officeConsents().deleteWithResponse("myRg", "myWorkspace", "04e5fd05-ff86-4b97-b8d2-1c20933cb46c",
            com.azure.core.util.Context.NONE);
    }
}
