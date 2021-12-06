Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupStatusRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.DataSourceType;

/** Samples for BackupStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureIaasVm/GetBackupStatus.json
     */
    /**
     * Sample code: Get Azure Virtual Machine Backup Status.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureVirtualMachineBackupStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupStatus()
            .getWithResponse(
                "southeastasia",
                new BackupStatusRequest()
                    .withResourceType(DataSourceType.VM)
                    .withResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Compute/VirtualMachines/testVm"),
                Context.NONE);
    }
}
```
