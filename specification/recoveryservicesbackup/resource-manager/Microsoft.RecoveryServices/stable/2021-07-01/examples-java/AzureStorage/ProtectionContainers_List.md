```java
import com.azure.core.util.Context;

/** Samples for BackupProtectionContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureStorage/ProtectionContainers_List.json
     */
    /**
     * Sample code: List Backup Protection Containers.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listBackupProtectionContainers(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupProtectionContainers()
            .list("testVault", "testRg", "backupManagementType eq 'AzureWorkload'", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
