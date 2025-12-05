
/**
 * Samples for Accounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Accounts_ListBySubscription.json
     */
    /**
     * Sample code: Accounts_ListBySubscription.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsListBySubscription(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().list(com.azure.core.util.Context.NONE);
    }
}
