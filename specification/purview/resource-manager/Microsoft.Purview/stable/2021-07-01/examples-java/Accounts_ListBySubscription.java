/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListBySubscription.json
     */
    /**
     * Sample code: Accounts_ListBySubscription.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsListBySubscription(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager.accounts().list(null, com.azure.core.util.Context.NONE);
    }
}
