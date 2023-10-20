/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/Accounts_Delete.json
     */
    /**
     * Sample code: Accounts_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().delete("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
