
/**
 * Samples for DdosProtectionPlans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/DdosProtectionPlanGet.json
     */
    /**
     * Sample code: Get DDoS protection plan.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDDoSProtectionPlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDdosProtectionPlans().getByResourceGroupWithResponse("rg1",
            "test-plan", com.azure.core.util.Context.NONE);
    }
}
