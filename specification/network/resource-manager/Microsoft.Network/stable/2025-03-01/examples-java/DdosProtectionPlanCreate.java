
import com.azure.resourcemanager.network.fluent.models.DdosProtectionPlanInner;

/**
 * Samples for DdosProtectionPlans CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/DdosProtectionPlanCreate.json
     */
    /**
     * Sample code: Create DDoS protection plan.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createDDoSProtectionPlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDdosProtectionPlans().createOrUpdate("rg1", "test-plan",
            new DdosProtectionPlanInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}
