/** Samples for IpPrefixes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpPrefixes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpPrefixes_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipPrefixesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipPrefixes().delete("rgIpPrefixLists", "example-ipPrefix", com.azure.core.util.Context.NONE);
    }
}
