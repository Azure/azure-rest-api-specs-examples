import com.azure.core.util.Context;

/** Samples for Routes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteTableRouteList.json
     */
    /**
     * Sample code: List routes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoutes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutes().list("rg1", "testrt", Context.NONE);
    }
}
