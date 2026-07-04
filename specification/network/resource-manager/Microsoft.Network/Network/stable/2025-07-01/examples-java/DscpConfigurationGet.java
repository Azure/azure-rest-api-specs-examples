
/**
 * Samples for DscpConfiguration GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DscpConfigurationGet.json
     */
    /**
     * Sample code: Get Dscp Configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getDscpConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDscpConfigurations().getByResourceGroupWithResponse("rg1", "mydscpConfig",
            com.azure.core.util.Context.NONE);
    }
}
