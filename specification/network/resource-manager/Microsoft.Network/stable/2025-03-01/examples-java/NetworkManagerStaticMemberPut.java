
import com.azure.resourcemanager.network.fluent.models.StaticMemberInner;

/**
 * Samples for StaticMembers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerStaticMemberPut
     * .json
     */
    /**
     * Sample code: StaticMemberPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticMemberPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getStaticMembers().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "testNetworkGroup", "testStaticMember",
            new StaticMemberInner().withResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/rg1/providers/Microsoft.Network/virtualnetworks/vnet1"),
            com.azure.core.util.Context.NONE);
    }
}
