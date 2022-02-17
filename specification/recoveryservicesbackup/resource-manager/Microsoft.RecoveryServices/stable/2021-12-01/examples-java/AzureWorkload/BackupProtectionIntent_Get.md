Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.3/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ProtectionIntent Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureWorkload/BackupProtectionIntent_Get.json
     */
    /**
     * Sample code: Get ProtectionIntent for an item.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectionIntentForAnItem(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionIntents()
            .getWithResponse("myVault", "myRG", "Azure", "249D9B07-D2EF-4202-AA64-65F35418564E", Context.NONE);
    }
}
```
