
/**
 * Samples for DdosProtectionPlans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosProtectionPlanGet.json
     */
    /**
     * Sample code: Get DDoS protection plan.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getDDoSProtectionPlan(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosProtectionPlans().getByResourceGroupWithResponse("rg1", "test-plan",
            com.azure.core.util.Context.NONE);
    }
}
