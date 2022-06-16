import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.DataSourceType;
import com.azure.resourcemanager.recoveryservicesbackup.models.PreValidateEnableBackupRequest;

/** Samples for ProtectionIntent Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ProtectionIntent_Validate.json
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
