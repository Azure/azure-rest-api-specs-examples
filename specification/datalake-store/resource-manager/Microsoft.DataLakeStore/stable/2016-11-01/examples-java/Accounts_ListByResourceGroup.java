/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists the Data Lake Store accounts within a specific resource group.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void listsTheDataLakeStoreAccountsWithinASpecificResourceGroup(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager
            .accounts()
            .listByResourceGroup(
                "contosorg",
                "test_filter",
                1,
                1,
                "test_select",
                "test_orderby",
                false,
                com.azure.core.util.Context.NONE);
    }
}
