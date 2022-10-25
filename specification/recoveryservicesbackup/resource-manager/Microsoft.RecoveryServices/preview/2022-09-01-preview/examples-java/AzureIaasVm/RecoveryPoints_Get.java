import com.azure.core.util.Context;

/** Samples for RecoveryPoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureIaasVm/RecoveryPoints_Get.json
     */
    /**
     * Sample code: Get Azure Vm Recovery Point Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureVmRecoveryPointDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .recoveryPoints()
            .getWithResponse(
                "rshvault",
                "rshhtestmdvmrg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
                "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
                "26083826328862",
                Context.NONE);
    }
}
