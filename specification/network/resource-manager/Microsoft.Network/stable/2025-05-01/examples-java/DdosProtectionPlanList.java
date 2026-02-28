
/**
 * Samples for DdosProtectionPlans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DdosProtectionPlanList.json
     */
    /**
     * Sample code: List DDoS protection plans in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDDoSProtectionPlansInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDdosProtectionPlans().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
