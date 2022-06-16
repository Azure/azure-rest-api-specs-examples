import com.azure.core.util.Context;

/** Samples for ExpressRoutePortsLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRoutePortsLocationList.json
     */
    /**
     * Sample code: ExpressRoutePortsLocationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortsLocationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePortsLocations().list(Context.NONE);
    }
}
