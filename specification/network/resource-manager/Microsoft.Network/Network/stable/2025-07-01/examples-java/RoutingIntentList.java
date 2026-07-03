
/**
 * Samples for RoutingIntent List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RoutingIntentList.json
     */
    /**
     * Sample code: RoutingIntentList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routingIntentList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRoutingIntents().list("rg1", "virtualHub1", com.azure.core.util.Context.NONE);
    }
}
