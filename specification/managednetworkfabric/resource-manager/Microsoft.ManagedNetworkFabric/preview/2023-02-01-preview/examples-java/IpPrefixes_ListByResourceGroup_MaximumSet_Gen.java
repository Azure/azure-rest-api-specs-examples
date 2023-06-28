/** Samples for IpPrefixes ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpPrefixes_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpPrefixes_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipPrefixesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipPrefixes().listByResourceGroup("resourcegroupname", com.azure.core.util.Context.NONE);
    }
}
