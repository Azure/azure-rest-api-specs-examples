Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.3/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupRequestResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMBackupRequest;

/** Samples for Backups Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/TriggerBackup_Post.json
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
```
