
/**
 * Samples for ProviderActions GetOverviewStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_GetOverviewStatus.json
     */
    /**
     * Sample code: Report_GetOverviewStatus.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportGetOverviewStatus(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().getOverviewStatusWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
