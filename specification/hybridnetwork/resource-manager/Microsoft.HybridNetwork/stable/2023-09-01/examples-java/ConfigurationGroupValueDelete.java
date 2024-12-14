
/**
 * Samples for ConfigurationGroupValues Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupValueDelete.json
     */
    /**
     * Sample code: Delete hybrid configuration group resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        deleteHybridConfigurationGroupResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupValues().delete("rg1", "testConfigurationGroupValue",
            com.azure.core.util.Context.NONE);
    }
}
