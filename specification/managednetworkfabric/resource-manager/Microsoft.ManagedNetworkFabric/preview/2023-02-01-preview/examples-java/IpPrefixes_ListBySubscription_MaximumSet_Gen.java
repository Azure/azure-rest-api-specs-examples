/** Samples for IpPrefixes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpPrefixes_ListBySubscription_MaximumSet_Gen.json
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
