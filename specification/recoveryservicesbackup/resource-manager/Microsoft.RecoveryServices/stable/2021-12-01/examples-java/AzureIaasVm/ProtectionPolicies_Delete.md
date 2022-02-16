Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.3/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ProtectionPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ProtectionPolicies_Delete.json
     */
    /**
     * Sample code: Delete Azure Vm Protection Policy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteAzureVmProtectionPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().delete("NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1", Context.NONE);
    }
}
```
