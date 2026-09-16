
/**
 * Samples for OfficeConsents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/officeConsents/GetOfficeConsents.json
     */
    /**
     * Sample code: Get all office consents.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllOfficeConsents(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.officeConsents().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
