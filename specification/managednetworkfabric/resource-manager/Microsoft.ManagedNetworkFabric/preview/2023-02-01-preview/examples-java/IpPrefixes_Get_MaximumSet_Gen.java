/** Samples for IpPrefixes GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpPrefixes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpPrefixes_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipPrefixesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .ipPrefixes()
            .getByResourceGroupWithResponse("resourcegroupname", "example-ipPrefix", com.azure.core.util.Context.NONE);
    }
}
