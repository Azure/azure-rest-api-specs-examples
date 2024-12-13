
/**
 * Samples for ConfigurationGroupValues GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupValueGet.json
     */
    /**
     * Sample code: Get hybrid configuration group.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        getHybridConfigurationGroup(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupValues().getByResourceGroupWithResponse("rg1", "testConfigurationGroupValue",
            com.azure.core.util.Context.NONE);
    }
}
