
/**
 * Samples for RecoveryPoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/RecoveryPoints_List.json
     */
    /**
     * Sample code: Get Protected Azure Vm Recovery Points.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedAzureVmRecoveryPoints(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.recoveryPoints().list("rshvault", "rshhtestmdvmrg", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
            "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall", null, com.azure.core.util.Context.NONE);
    }
}
