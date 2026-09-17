
/**
 * Samples for VirtualApplianceSkus Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetworkVirtualApplianceSkuGet.json
     */
    /**
     * Sample code: NetworkVirtualApplianceSkuGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void networkVirtualApplianceSkuGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualApplianceSkus().getWithResponse("ciscoSdwan",
            com.azure.core.util.Context.NONE);
    }
}
