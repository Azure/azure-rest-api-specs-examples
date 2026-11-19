
/**
 * Samples for NetworkVirtualAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/NetworkVirtualApplianceVhubSourceIPProtocolGet.json
     */
    /**
     * Sample code: Get a managed NVA in a virtual hub with SourceIPProtocol load distribution.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAManagedNVAInAVirtualHubWithSourceIPProtocolLoadDistribution(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().getByResourceGroupWithResponse("rg1", "nva-vhub", null,
            com.azure.core.util.Context.NONE);
    }
}
