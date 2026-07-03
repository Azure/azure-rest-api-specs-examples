
/**
 * Samples for NetworkSecurityPerimeterLoggingConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLoggingConfigurationGet.json
     */
    /**
     * Sample code: NspLoggingConfigurationGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLoggingConfigurationGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLoggingConfigurations().getWithResponse("rg1", "nsp1",
            "instance", com.azure.core.util.Context.NONE);
    }
}
