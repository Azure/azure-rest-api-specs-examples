
/**
 * Samples for IpPrefixes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpPrefixes_Get.json
     */
    /**
     * Sample code: IpPrefixes_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void
        ipPrefixesGetMaximumSetGen(com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipPrefixes().getByResourceGroupWithResponse("example-rg", "example-ipPrefix",
            com.azure.core.util.Context.NONE);
    }
}
