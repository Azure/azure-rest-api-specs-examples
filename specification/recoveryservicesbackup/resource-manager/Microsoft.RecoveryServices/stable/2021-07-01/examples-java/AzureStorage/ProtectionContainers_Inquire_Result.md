Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ProtectionContainerOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureStorage/ProtectionContainers_Inquire_Result.json
     */
    /**
     * Sample code: Get Azure Storage Protection Container Operation Result.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureStorageProtectionContainerOperationResult(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainerOperationResults()
            .getWithResponse(
                "testvault",
                "test-rg",
                "Azure",
                "VMAppContainer;Compute;testRG;testSQL",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
```
