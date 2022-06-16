import com.azure.core.util.Context;

/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/Accounts_List.json
     */
    /**
     * Sample code: Accounts_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().list(Context.NONE);
    }
}
