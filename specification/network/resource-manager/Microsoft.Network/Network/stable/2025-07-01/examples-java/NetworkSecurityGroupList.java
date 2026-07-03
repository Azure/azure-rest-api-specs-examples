
/**
 * Samples for NetworkSecurityGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupList.json
     */
    /**
     * Sample code: List network security groups in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listNetworkSecurityGroupsInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityGroups().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
