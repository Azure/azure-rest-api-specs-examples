
/**
 * Samples for IpPrefixes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpPrefixes_ListByResourceGroup.json
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
