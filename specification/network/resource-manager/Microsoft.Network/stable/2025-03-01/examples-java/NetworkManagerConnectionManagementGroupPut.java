
import com.azure.resourcemanager.network.fluent.models.NetworkManagerConnectionInner;

/**
 * Samples for ManagementGroupNetworkManagerConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerConnectionManagementGroupPut.json
     */
    /**
     * Sample code: Create or Update Management Group Network Manager Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateManagementGroupNetworkManagerConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getManagementGroupNetworkManagerConnections()
            .createOrUpdateWithResponse("managementGroupA", "TestNMConnection",
                new NetworkManagerConnectionInner().withNetworkManagerId(
                    "/subscriptions/subscriptionC/resourceGroup/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager"),
                com.azure.core.util.Context.NONE);
    }
}
