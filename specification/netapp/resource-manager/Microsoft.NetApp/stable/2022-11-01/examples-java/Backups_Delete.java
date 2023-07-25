/** Samples for Backups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Backups_Delete.json
     */
    /**
     * Sample code: Backups_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .backups()
            .delete(
                "resourceGroup",
                "accountName",
                "poolName",
                "volumeName",
                "backupName",
                com.azure.core.util.Context.NONE);
    }
}
