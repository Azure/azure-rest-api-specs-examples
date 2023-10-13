/** Samples for RecoveryPoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/RecoveryPoints_List.json
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
                com.azure.core.util.Context.NONE);
    }
}
