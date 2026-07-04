
/**
 * Samples for IpGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpGroupsListBySubscription.json
     */
    /**
     * Sample code: List_IpGroups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listIpGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpGroups().list(com.azure.core.util.Context.NONE);
    }
}
