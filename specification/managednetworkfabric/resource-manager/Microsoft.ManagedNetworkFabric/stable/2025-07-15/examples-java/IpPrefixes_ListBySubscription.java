
/**
 * Samples for IpPrefixes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpPrefixes_ListBySubscription.json
     */
    /**
     * Sample code: IpPrefixes_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipPrefixesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipPrefixes().list(com.azure.core.util.Context.NONE);
    }
}
