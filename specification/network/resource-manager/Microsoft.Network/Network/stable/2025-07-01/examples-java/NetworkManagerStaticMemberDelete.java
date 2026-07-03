
/**
 * Samples for StaticMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerStaticMemberDelete.json
     */
    /**
     * Sample code: StaticMembersDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticMembersDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticMembers().deleteWithResponse("SampleRG", "TestNM", "testNetworkGroup",
            "testStaticMember", com.azure.core.util.Context.NONE);
    }
}
