
/**
 * Samples for PrivateZones ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/PrivateZoneListInResourceGroup.json
     */
    /**
     * Sample code: GET Private DNS Zone by Resource Group.
     * 
     * @param manager Entry point to PrivateDnsZoneManager.
     */
    public static void
        gETPrivateDNSZoneByResourceGroup(com.azure.resourcemanager.privatedns.PrivateDnsZoneManager manager) {
        manager.serviceClient().getPrivateZones().listByResourceGroup("resourceGroup1", null,
            com.azure.core.util.Context.NONE);
    }
}
