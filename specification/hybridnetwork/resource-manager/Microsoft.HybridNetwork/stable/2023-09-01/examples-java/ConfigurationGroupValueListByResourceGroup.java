
/**
 * Samples for ConfigurationGroupValues ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupValueListByResourceGroup.json
     */
    /**
     * Sample code: List all hybrid network configurationGroupValues in a subscription.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllHybridNetworkConfigurationGroupValuesInASubscription(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupValues().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
