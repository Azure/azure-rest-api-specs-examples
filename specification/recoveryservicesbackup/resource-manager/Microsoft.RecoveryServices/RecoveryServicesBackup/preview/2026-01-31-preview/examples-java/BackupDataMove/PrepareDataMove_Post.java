
import com.azure.resourcemanager.recoveryservicesbackup.models.DataMoveLevel;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrepareDataMoveRequest;

/**
 * Samples for ResourceProvider BmsPrepareDataMove.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/BackupDataMove/PrepareDataMove_Post.json
     */
    /**
     * Sample code: Prepare Data Move.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        prepareDataMove(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.resourceProviders().bmsPrepareDataMove("source-rsv", "sourceRG",
            new PrepareDataMoveRequest().withTargetResourceId(
                "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/targetRG/providers/Microsoft.RecoveryServices/vaults/target-rsv")
                .withTargetRegion("USGov Virginia").withDataMoveLevel(DataMoveLevel.VAULT),
            com.azure.core.util.Context.NONE);
    }
}
