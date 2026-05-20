
/**
 * Samples for ProviderActions Onboard.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Onboard.json
     */
    /**
     * Sample code: Onboard.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        onboard(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().onboard(null, com.azure.core.util.Context.NONE);
    }
}
