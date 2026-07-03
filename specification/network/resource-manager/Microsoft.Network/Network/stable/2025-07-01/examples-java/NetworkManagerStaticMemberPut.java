
import com.azure.resourcemanager.network.fluent.models.StaticMemberInner;

/**
 * Samples for StaticMembers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerStaticMemberPut.json
     */
    /**
     * Sample code: StaticMemberPut.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticMemberPut(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticMembers().createOrUpdateWithResponse("rg1", "testNetworkManager",
            "testNetworkGroup", "testStaticMember",
            new StaticMemberInner().withResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualnetworks/vnet1"),
            com.azure.core.util.Context.NONE);
    }
}
