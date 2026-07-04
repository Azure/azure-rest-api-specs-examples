
/**
 * Samples for StaticMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerStaticMemberGet.json
     */
    /**
     * Sample code: StaticMembersGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticMembersGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticMembers().getWithResponse("rg1", "testNetworkManager", "testNetworkGroup",
            "testStaticMember", com.azure.core.util.Context.NONE);
    }
}
