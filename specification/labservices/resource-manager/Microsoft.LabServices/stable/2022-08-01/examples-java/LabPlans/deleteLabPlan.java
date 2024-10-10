
/**
 * Samples for LabPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/
     * deleteLabPlan.json
     */
    /**
     * Sample code: deleteLabPlan.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void deleteLabPlan(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labPlans().delete("testrg123", "testlabplan", com.azure.core.util.Context.NONE);
    }
}
