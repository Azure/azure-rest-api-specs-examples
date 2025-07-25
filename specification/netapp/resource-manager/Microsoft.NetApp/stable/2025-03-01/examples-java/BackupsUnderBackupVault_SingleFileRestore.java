
import com.azure.resourcemanager.netapp.models.BackupRestoreFiles;
import java.util.Arrays;

/**
 * Samples for BackupsUnderBackupVault RestoreFiles.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/
     * BackupsUnderBackupVault_SingleFileRestore.json
     */
    /**
     * Sample code: Backups_SingleFileRestore.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsSingleFileRestore(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupsUnderBackupVaults().restoreFiles("myRG", "account1", "backupVault1", "backup1",
            new BackupRestoreFiles().withFileList(Arrays.asList("/dir1/customer1.db", "/dir1/customer2.db"))
                .withDestinationVolumeId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1"),
            com.azure.core.util.Context.NONE);
    }
}
