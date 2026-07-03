
/**
 * Samples for ApplicationSecurityGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationSecurityGroupList.json
     */
    /**
     * Sample code: List load balancers in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listLoadBalancersInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationSecurityGroups().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
