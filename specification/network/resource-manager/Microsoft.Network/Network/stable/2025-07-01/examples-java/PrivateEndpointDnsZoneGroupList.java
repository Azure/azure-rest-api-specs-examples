
/**
 * Samples for PrivateDnsZoneGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointDnsZoneGroupList.json
     */
    /**
     * Sample code: List private endpoints in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPrivateEndpointsInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateDnsZoneGroups().list("testPe", "rg1", com.azure.core.util.Context.NONE);
    }
}
