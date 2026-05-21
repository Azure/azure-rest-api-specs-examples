
/**
 * Samples for ProviderActions GetCollectionCount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_GetCollectionCount.json
     */
    /**
     * Sample code: Report_GetCollectionCount.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportGetCollectionCount(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().getCollectionCountWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
