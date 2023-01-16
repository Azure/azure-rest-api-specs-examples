/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_List.json
     */
    /**
     * Sample code: Lists the Data Lake Store accounts within the subscription.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void listsTheDataLakeStoreAccountsWithinTheSubscription(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager
            .accounts()
            .list("test_filter", 1, 1, "test_select", "test_orderby", false, com.azure.core.util.Context.NONE);
    }
}
