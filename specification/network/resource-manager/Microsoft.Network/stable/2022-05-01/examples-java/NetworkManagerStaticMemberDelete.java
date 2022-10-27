import com.azure.core.util.Context;

/** Samples for StaticMembers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerStaticMemberDelete.json
     */
    /**
     * Sample code: StaticMembersDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticMembersDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getStaticMembers()
            .deleteWithResponse("SampleRG", "TestNM", "testNetworkGroup", "testStaticMember", Context.NONE);
    }
}
