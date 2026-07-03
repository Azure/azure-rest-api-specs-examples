
/**
 * Samples for RouteTables List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableListAll.json
     */
    /**
     * Sample code: List all route tables.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllRouteTables(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().list(com.azure.core.util.Context.NONE);
    }
}
