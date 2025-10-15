
/**
 * Samples for Accounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/
     * Accounts_ListBySubscription.json
     */
    /**
     * Sample code: Accounts_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().list(com.azure.core.util.Context.NONE);
    }
}
