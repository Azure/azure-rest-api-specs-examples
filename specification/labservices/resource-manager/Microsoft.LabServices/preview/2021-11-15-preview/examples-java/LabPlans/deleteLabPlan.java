import com.azure.core.util.Context;

/** Samples for LabPlans Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/deleteLabPlan.json
     */
    /**
     * Sample code: deleteLabPlan.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void deleteLabPlan(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labPlans().delete("testrg123", "testlabplan", Context.NONE);
    }
}
