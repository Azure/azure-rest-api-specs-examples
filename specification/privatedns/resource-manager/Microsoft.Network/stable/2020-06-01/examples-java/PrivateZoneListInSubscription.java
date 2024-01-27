
/** Samples for PrivateZones List. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/
     * PrivateZoneListInSubscription.json
     */
    /**
     * Sample code: GET Private DNS Zone by Subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETPrivateDNSZoneBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getPrivateZones().list(null,
            com.azure.core.util.Context.NONE);
    }
}
