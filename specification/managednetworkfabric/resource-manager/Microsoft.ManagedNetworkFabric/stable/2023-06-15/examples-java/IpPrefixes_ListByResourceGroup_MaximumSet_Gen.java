/** Samples for IpPrefixes ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpPrefixes_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpPrefixes_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipPrefixesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipPrefixes().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
