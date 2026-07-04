
/**
 * Samples for PrivateDnsZoneGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointDnsZoneGroupGet.json
     */
    /**
     * Sample code: Get private dns zone group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPrivateDnsZoneGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateDnsZoneGroups().getWithResponse("rg1", "testPe", "testPdnsgroup",
            com.azure.core.util.Context.NONE);
    }
}
