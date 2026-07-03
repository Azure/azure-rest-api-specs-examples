
/**
 * Samples for StaticMembers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerStaticMemberList.json
     */
    /**
     * Sample code: StaticMembersList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticMembersList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticMembers().list("rg1", "testNetworkManager", "testNetworkGroup", null, null,
            com.azure.core.util.Context.NONE);
    }
}
