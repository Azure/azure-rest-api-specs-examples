
/**
 * Samples for DdosCustomPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DdosCustomPolicyList.json
     */
    /**
     * Sample code: List DDoS custom policies in a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listDDoSCustomPoliciesInAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosCustomPolicies().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
