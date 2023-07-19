/** Samples for IpPrefixes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpPrefixes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpPrefixes_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipPrefixesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipPrefixes().delete("example-rg", "example-ipPrefix", com.azure.core.util.Context.NONE);
    }
}
