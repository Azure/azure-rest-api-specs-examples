
/**
 * Samples for IpGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpGroupsGet.json
     */
    /**
     * Sample code: Get_IpGroups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getIpGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpGroups().getByResourceGroupWithResponse("myResourceGroup", "ipGroups1", null,
            com.azure.core.util.Context.NONE);
    }
}
