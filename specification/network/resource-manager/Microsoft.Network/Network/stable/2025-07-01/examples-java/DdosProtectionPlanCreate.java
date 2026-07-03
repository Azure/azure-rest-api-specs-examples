
import com.azure.resourcemanager.network.fluent.models.DdosProtectionPlanInner;

/**
 * Samples for DdosProtectionPlans CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosProtectionPlanCreate.json
     */
    /**
     * Sample code: Create DDoS protection plan.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createDDoSProtectionPlan(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosProtectionPlans().createOrUpdate("rg1", "test-plan",
            new DdosProtectionPlanInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}
