import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuits List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitListBySubscription.json
     */
    /**
     * Sample code: List ExpressRouteCircuits in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressRouteCircuitsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuits().list(Context.NONE);
    }
}
