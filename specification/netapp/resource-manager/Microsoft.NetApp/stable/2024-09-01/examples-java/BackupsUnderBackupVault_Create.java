
/**
 * Samples for Backups Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/BackupsUnderBackupVault_Create.
     * json
     */
    /**
     * Sample code: BackupsUnderBackupVault_Create.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUnderBackupVaultCreate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().define("backup1").withExistingBackupVault("myRG", "account1", "backupVault1")
            .withVolumeResourceId(
                "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPool/pool1/volumes/volume1")
            .withLabel("myLabel").create();
    }
}
