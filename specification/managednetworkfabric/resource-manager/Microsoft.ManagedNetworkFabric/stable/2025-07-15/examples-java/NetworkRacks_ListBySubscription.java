
/**
 * Samples for NetworkRacks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkRacks_ListBySubscription.json
     */
    /**
     * Sample code: NetworkRacks_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkRacksListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkRacks().list(com.azure.core.util.Context.NONE);
    }
}
