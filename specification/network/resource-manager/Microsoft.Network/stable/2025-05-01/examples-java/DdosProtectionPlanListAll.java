
/**
 * Samples for DdosProtectionPlans List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DdosProtectionPlanListAll.
     * json
     */
    /**
     * Sample code: List all DDoS protection plans.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllDDoSProtectionPlans(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDdosProtectionPlans().list(com.azure.core.util.Context.NONE);
    }
}
