
/**
 * Samples for DdosProtectionPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosProtectionPlanDelete.json
     */
    /**
     * Sample code: Delete DDoS protection plan.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteDDoSProtectionPlan(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosProtectionPlans().delete("rg1", "test-plan", com.azure.core.util.Context.NONE);
    }
}
