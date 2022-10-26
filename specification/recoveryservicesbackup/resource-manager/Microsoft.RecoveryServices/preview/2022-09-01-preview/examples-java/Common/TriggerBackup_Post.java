import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupRequestResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMBackupRequest;

/** Samples for Backups Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/TriggerBackup_Post.json
     */
    /**
     * Sample code: Trigger Backup.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void triggerBackup(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backups()
            .triggerWithResponse(
                "linuxRsVault",
                "linuxRsVaultRG",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;testrg;v1win2012r",
                "VM;iaasvmcontainerv2;testrg;v1win2012r",
                new BackupRequestResource().withProperties(new IaasVMBackupRequest()),
                Context.NONE);
    }
}
