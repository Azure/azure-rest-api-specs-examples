
/**
 * Samples for DscpConfiguration ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DscpConfigurationList.json
     */
    /**
     * Sample code: Get Dscp Configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getDscpConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDscpConfigurations().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
