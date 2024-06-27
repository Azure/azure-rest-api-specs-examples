
/**
 * Samples for Evidence CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/Evidence_CreateOrUpdate.json
     */
    /**
     * Sample code: Evidence_CreateOrUpdate.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void evidenceCreateOrUpdate(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.evidences().createOrUpdateWithResponse("testReportName", "evidence1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
