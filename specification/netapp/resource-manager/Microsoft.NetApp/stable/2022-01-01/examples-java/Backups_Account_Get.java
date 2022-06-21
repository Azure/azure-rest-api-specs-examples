import com.azure.core.util.Context;

/** Samples for AccountBackups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Backups_Account_Get.json
     */
    /**
     * Sample code: AccountBackups_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountBackupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accountBackups().getWithResponse("myRG", "account1", "backup1", Context.NONE);
    }
}
