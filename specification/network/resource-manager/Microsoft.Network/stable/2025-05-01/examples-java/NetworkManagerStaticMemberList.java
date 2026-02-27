
/**
 * Samples for StaticMembers List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerStaticMemberList.json
     */
    /**
     * Sample code: StaticMembersList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticMembersList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getStaticMembers().list("rg1", "testNetworkManager",
            "testNetworkGroup", null, null, com.azure.core.util.Context.NONE);
    }
}
