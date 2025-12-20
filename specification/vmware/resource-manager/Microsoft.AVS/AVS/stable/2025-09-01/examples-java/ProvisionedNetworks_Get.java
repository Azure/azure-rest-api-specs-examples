
/**
 * Samples for ProvisionedNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ProvisionedNetworks_Get.json
     */
    /**
     * Sample code: ProvisionedNetworks_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void provisionedNetworksGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.provisionedNetworks().getWithResponse("group1", "cloud1", "vsan", com.azure.core.util.Context.NONE);
    }
}
