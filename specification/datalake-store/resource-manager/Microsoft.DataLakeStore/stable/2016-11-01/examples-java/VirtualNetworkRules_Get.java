/** Samples for VirtualNetworkRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Get.json
     */
    /**
     * Sample code: Gets the specified Data Lake Store virtual network rule.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void getsTheSpecifiedDataLakeStoreVirtualNetworkRule(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager
            .virtualNetworkRules()
            .getWithResponse(
                "contosorg", "contosoadla", "test_virtual_network_rules_name", com.azure.core.util.Context.NONE);
    }
}
