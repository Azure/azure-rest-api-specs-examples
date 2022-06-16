import com.azure.core.util.Context;

/** Samples for ProtectedItemOperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/ProtectedItemOperationStatus.json
     */
    /**
     * Sample code: Get Operation Status of Protected Vm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationStatusOfProtectedVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectedItemOperationStatuses()
            .getWithResponse(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
