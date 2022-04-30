Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.4/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BackupEngines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Dpm/BackupEngines_Get.json
     */
    /**
     * Sample code: Get Dpm/AzureBackupServer/Lajolla Backup Engine Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getDpmAzureBackupServerLajollaBackupEngineDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupEngines().getWithResponse("testVault", "testRG", "testServer", null, null, Context.NONE);
    }
}
```
