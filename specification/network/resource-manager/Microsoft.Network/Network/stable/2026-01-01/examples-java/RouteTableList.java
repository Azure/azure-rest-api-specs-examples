
/**
 * Samples for RouteTables ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/RouteTableList.json
     */
    /**
     * Sample code: List route tables in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listRouteTablesInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
