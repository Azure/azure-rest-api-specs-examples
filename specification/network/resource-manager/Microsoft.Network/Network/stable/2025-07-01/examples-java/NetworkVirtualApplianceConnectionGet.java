
/**
 * Samples for NetworkVirtualApplianceConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceConnectionGet.json
     */
    /**
     * Sample code: NetworkVirtualApplianceConnectionGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkVirtualApplianceConnectionGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualApplianceConnections().getWithResponse("rg1", "nva1", "connection1",
            com.azure.core.util.Context.NONE);
    }
}
