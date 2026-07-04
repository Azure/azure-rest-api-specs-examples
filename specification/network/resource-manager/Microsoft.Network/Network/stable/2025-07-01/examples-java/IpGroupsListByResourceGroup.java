
/**
 * Samples for IpGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpGroupsListByResourceGroup.json
     */
    /**
     * Sample code: ListByResourceGroup_IpGroups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listByResourceGroupIpGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpGroups().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
