
/**
 * Samples for IpPrefixes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpPrefixes_Delete.json
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
