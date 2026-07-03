
/**
 * Samples for PrivateDnsZoneGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointDnsZoneGroupDelete.json
     */
    /**
     * Sample code: Delete private dns zone group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePrivateDnsZoneGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateDnsZoneGroups().delete("rg1", "testPe", "testPdnsgroup",
            com.azure.core.util.Context.NONE);
    }
}
