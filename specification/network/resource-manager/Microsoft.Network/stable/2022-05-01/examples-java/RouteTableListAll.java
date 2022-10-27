import com.azure.core.util.Context;

/** Samples for RouteTables List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/RouteTableListAll.json
     */
    /**
     * Sample code: List all route tables.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllRouteTables(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteTables().list(Context.NONE);
    }
}
