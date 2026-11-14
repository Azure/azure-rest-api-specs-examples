
/**
 * Samples for PrivateZones List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/PrivateZoneListInSubscription.json
     */
    /**
     * Sample code: GET Private DNS Zone by Subscription.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZoneBySubscription(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getPrivateZones().list(null, com.azure.core.util.Context.NONE);
    }
}
