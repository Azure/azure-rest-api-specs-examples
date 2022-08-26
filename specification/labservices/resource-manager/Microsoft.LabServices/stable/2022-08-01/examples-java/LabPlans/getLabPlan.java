import com.azure.core.util.Context;

/** Samples for LabPlans GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/getLabPlan.json
     */
    /**
     * Sample code: getLabPlan.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getLabPlan(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labPlans().getByResourceGroupWithResponse("testrg123", "testlabplan", Context.NONE);
    }
}
