import com.azure.core.util.Context;

/** Samples for RoutingIntent List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/RoutingIntentList.json
     */
    /**
     * Sample code: RoutingIntentList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routingIntentList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRoutingIntents().list("rg1", "virtualHub1", Context.NONE);
    }
}
