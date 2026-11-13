
/**
 * Samples for PrivateZones GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/PrivateZoneGet.json
     */
    /**
     * Sample code: GET Private DNS Zone.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void gETPrivateDNSZone(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getPrivateZones().getByResourceGroupWithResponse("resourceGroup1", "privatezone1.com",
            com.azure.core.util.Context.NONE);
    }
}
