
/**
 * Samples for DdosProtectionPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/DdosProtectionPlanDelete.json
     */
    /**
     * Sample code: Delete DDoS protection plan.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteDDoSProtectionPlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDdosProtectionPlans().delete("rg1", "test-plan",
            com.azure.core.util.Context.NONE);
    }
}
