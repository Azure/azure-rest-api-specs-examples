
/**
 * Samples for NetworkRacks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkRacks_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkRacks_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkRacksListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkRacks().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
