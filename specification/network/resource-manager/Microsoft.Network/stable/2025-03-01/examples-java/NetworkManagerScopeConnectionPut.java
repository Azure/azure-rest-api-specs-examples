
import com.azure.resourcemanager.network.fluent.models.ScopeConnectionInner;

/**
 * Samples for ScopeConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerScopeConnectionPut.json
     */
    /**
     * Sample code: Create or Update Network Manager Scope Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateNetworkManagerScopeConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getScopeConnections().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "TestScopeConnection",
            new ScopeConnectionInner().withTenantId("6babcaad-604b-40ac-a9d7-9fd97c0b779f")
                .withResourceId("subscriptions/f0dc2b34-dfad-40e4-83e0-2309fed8d00b").withDescription(
                    "This is a scope connection to a cross tenant subscription."),
            com.azure.core.util.Context.NONE);
    }
}
