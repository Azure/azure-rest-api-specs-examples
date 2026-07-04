
import com.azure.resourcemanager.network.fluent.models.NetworkManagerConnectionInner;

/**
 * Samples for SubscriptionNetworkManagerConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectionSubscriptionPut.json
     */
    /**
     * Sample code: Create or Update Subscription Network Manager Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createOrUpdateSubscriptionNetworkManagerConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubscriptionNetworkManagerConnections()
            .createOrUpdateWithResponse("TestNMConnection", new NetworkManagerConnectionInner().withNetworkManagerId(
                "/subscriptions/22222222-2222-2222-2222-222222222222/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager"),
                com.azure.core.util.Context.NONE);
    }
}
