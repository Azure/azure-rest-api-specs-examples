Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.3/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BackupOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/ProtectedItem_Delete_OperationResult.json
     */
    /**
     * Sample code: Get Result for Protected Item Delete Operation.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getResultForProtectedItemDeleteOperation(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupOperationResults()
            .getWithResponse(
                "PySDKBackupTestRsVault",
                "PythonSDKBackupTestRg",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
```
