
/**
 * Samples for Report NestedResourceCheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/Report_NestedResourceCheckNameAvailability_Report_Webhook_Check_Name_Availability.json
     */
    /**
     * Sample code: Report_WebhookCheckNameAvailability.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportWebhookCheckNameAvailability(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().nestedResourceCheckNameAvailabilityWithResponse("reportABC", null,
            com.azure.core.util.Context.NONE);
    }
}
