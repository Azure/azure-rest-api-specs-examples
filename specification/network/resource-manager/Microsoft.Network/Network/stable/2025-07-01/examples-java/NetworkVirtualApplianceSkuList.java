
/**
 * Samples for VirtualApplianceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceSkuList.json
     */
    /**
     * Sample code: NetworkVirtualApplianceSkuListResult.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkVirtualApplianceSkuListResult(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualApplianceSkus().list(com.azure.core.util.Context.NONE);
    }
}
