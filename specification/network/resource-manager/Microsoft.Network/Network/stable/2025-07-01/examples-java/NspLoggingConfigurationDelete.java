
/**
 * Samples for NetworkSecurityPerimeterLoggingConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLoggingConfigurationDelete.json
     */
    /**
     * Sample code: NspLoggingConfigurationDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLoggingConfigurationDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLoggingConfigurations().deleteWithResponse("rg1", "nsp1",
            "instance", com.azure.core.util.Context.NONE);
    }
}
