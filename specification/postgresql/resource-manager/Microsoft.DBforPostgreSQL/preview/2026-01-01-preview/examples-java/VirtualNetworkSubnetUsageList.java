
import com.azure.resourcemanager.postgresqlflexibleserver.models.VirtualNetworkSubnetUsageParameter;

/**
 * Samples for VirtualNetworkSubnetUsage List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/VirtualNetworkSubnetUsageList.json
     */
    /**
     * Sample code: List the virtual network subnet usage for a given virtual network.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheVirtualNetworkSubnetUsageForAGivenVirtualNetwork(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.virtualNetworkSubnetUsages().listWithResponse("eastus",
            new VirtualNetworkSubnetUsageParameter().withVirtualNetworkArmResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/virtualNetworks/examplevirtualnetwork"),
            com.azure.core.util.Context.NONE);
    }
}
