import com.azure.core.util.Context;

/** Samples for AccountBackups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/Backups_Account_List.json
     */
    /**
     * Sample code: AccountBackups_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountBackupsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accountBackups().list("myRG", "account1", Context.NONE);
    }
}
