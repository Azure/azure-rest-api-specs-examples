Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.4/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BackupWorkloadItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureWorkload/BackupWorkloadItems_List.json
     */
    /**
     * Sample code: List Workload Items in Container.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listWorkloadItemsInContainer(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupWorkloadItems()
            .list(
                "suchandr-seacan-rsv",
                "testRg",
                "Azure",
                "VMAppContainer;Compute;bvtdtestag;sqlserver-1",
                "backupManagementType eq 'AzureWorkload'",
                null,
                Context.NONE);
    }
}
```
