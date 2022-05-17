Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.5/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.DataSourceType;
import com.azure.resourcemanager.recoveryservicesbackup.models.PreValidateEnableBackupRequest;

/** Samples for ProtectionIntent Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/ProtectionIntent_Validate.json
     */
    /**
     * Sample code: Validate Enable Protection on Azure Vm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void validateEnableProtectionOnAzureVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionIntents()
            .validateWithResponse(
                "southeastasia",
                new PreValidateEnableBackupRequest()
                    .withResourceType(DataSourceType.VM)
                    .withResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/arunaupgrade/providers/Microsoft.Compute/VirtualMachines/upgrade1")
                    .withVaultId(
                        "/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/Vaults/myVault")
                    .withProperties(""),
                Context.NONE);
    }
}
```
