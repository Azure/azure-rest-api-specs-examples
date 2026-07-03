
/**
 * Samples for IpGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpGroupsDelete.json
     */
    /**
     * Sample code: Delete_IpGroups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteIpGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpGroups().delete("myResourceGroup", "ipGroups1", com.azure.core.util.Context.NONE);
    }
}
