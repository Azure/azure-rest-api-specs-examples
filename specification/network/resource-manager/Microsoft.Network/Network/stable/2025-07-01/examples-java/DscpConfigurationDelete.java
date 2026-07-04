
/**
 * Samples for DscpConfiguration Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DscpConfigurationDelete.json
     */
    /**
     * Sample code: Delete DSCP Configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteDSCPConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDscpConfigurations().delete("rg1", "mydscpConfig", com.azure.core.util.Context.NONE);
    }
}
