import com.azure.resourcemanager.recoveryservicessiterecovery.models.MigrationItem;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateMigrationItemInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VMwareCbtUpdateMigrationItemInput;

/** Samples for ReplicationMigrationItems Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationMigrationItems_Update.json
     */
    /**
     * Sample code: Updates migration item.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updatesMigrationItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        MigrationItem resource =
            manager
                .replicationMigrationItems()
                .getWithResponse(
                    "migrationvault",
                    "resourcegroup1",
                    "vmwarefabric1",
                    "vmwareContainer1",
                    "virtualmachine1",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new UpdateMigrationItemInputProperties()
                    .withProviderSpecificDetails(new VMwareCbtUpdateMigrationItemInput()))
            .apply();
    }
}
