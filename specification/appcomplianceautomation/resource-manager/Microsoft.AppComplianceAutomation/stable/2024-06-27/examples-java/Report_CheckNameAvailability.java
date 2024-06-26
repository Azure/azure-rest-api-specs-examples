
/**
 * Samples for ProviderActions CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/Report_CheckNameAvailability.json
     */
    /**
     * Sample code: Report_CheckNameAvailability.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportCheckNameAvailability(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().checkNameAvailabilityWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
