
/**
 * Samples for PrivateZones Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/PrivateZoneDelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void dELETEPrivateDNSZone(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getPrivateZones().delete("resourceGroup1", "privatezone1.com", null,
            com.azure.core.util.Context.NONE);
    }
}
