
import com.azure.resourcemanager.mysqlflexibleserver.models.VirtualNetworkSubnetUsageParameter;

/**
 * Samples for CheckVirtualNetworkSubnetUsage Execute.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/
     * CheckVirtualNetworkSubnetUsage.json
     */
    /**
     * Sample code: CheckVirtualNetworkSubnetUsage.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void
        checkVirtualNetworkSubnetUsage(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.checkVirtualNetworkSubnetUsages().executeWithResponse("WestUS",
            new VirtualNetworkSubnetUsageParameter().withVirtualNetworkResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/testvnet"),
            com.azure.core.util.Context.NONE);
    }
}
