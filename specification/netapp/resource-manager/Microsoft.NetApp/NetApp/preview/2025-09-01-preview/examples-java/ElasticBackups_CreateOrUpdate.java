
import com.azure.resourcemanager.netapp.models.ElasticBackupProperties;
import com.azure.resourcemanager.netapp.models.SnapshotUsage;

/**
 * Samples for ElasticBackups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackups_CreateOrUpdate.json
     */
    /**
     * Sample code: ElasticBackups_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackups().define("backup1").withExistingElasticBackupVault("myRG", "account1", "backupVault1")
            .withProperties(new ElasticBackupProperties().withLabel("myLabel").withElasticVolumeResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1")
                .withSnapshotUsage(SnapshotUsage.USE_EXISTING_SNAPSHOT).withElasticSnapshotResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1/elasticSnapshots/snap1"))
            .create();
    }
}
