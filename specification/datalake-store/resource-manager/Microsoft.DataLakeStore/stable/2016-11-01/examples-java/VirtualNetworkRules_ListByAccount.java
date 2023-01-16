/** Samples for VirtualNetworkRules ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_ListByAccount.json
     */
    /**
     * Sample code: Lists the Data Lake Store virtual network rules within the specified Data Lake Store account.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void listsTheDataLakeStoreVirtualNetworkRulesWithinTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.virtualNetworkRules().listByAccount("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}
