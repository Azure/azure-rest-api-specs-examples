/** Samples for AccountBackups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Backups_Account_Delete.json
     */
    /**
     * Sample code: AccountBackups_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountBackupsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accountBackups().delete("resourceGroup", "accountName", "backupName", com.azure.core.util.Context.NONE);
    }
}
