Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.3/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BackupUsageSummaries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/BackupProtectedItem_UsageSummary_Get.json
     */
    /**
     * Sample code: Get Protected Items Usages Summary.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedItemsUsagesSummary(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupUsageSummaries()
            .list("testVault", "testRG", "type eq 'BackupProtectedItemCountSummary'", null, Context.NONE);
    }
}
```
