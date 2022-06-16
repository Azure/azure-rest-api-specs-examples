import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresqlflexibleserver.models.VirtualNetworkSubnetUsageParameter;

/** Samples for VirtualNetworkSubnetUsage Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/VirtualNetworkSubnetUsage.json
     */
    /**
     * Sample code: VirtualNetworkSubnetUsageList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void virtualNetworkSubnetUsageList(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .virtualNetworkSubnetUsages()
            .executeWithResponse(
                "westus",
                new VirtualNetworkSubnetUsageParameter()
                    .withVirtualNetworkArmResourceId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/testvnet"),
                Context.NONE);
    }
}
