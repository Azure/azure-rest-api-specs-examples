
/**
 * Samples for DdosProtectionPlans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosProtectionPlanListAll.json
     */
    /**
     * Sample code: List all DDoS protection plans.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllDDoSProtectionPlans(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosProtectionPlans().list(com.azure.core.util.Context.NONE);
    }
}
