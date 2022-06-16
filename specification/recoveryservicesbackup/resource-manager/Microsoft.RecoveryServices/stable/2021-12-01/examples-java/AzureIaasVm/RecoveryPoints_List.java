import com.azure.core.util.Context;

/** Samples for RecoveryPoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/RecoveryPoints_List.json
     */
    /**
     * Sample code: Get Protected Azure Vm Recovery Points.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedAzureVmRecoveryPoints(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .recoveryPoints()
            .list(
                "rshvault",
                "rshhtestmdvmrg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
                "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
                null,
                Context.NONE);
    }
}
