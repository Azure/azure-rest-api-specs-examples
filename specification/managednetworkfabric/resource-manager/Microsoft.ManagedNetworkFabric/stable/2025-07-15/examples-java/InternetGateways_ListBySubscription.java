
/**
 * Samples for InternetGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternetGateways_ListBySubscription.json
     */
    /**
     * Sample code: InternetGateways_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewaysListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGateways().list(com.azure.core.util.Context.NONE);
    }
}
