import com.azure.core.util.Context;

/** Samples for Zones List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListZonesBySubscription.json
     */
    /**
     * Sample code: List zones by subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listZonesBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getZones().list(null, Context.NONE);
    }
}
