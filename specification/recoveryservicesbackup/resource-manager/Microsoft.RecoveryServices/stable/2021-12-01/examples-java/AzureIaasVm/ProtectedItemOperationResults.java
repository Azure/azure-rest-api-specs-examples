import com.azure.core.util.Context;

/** Samples for ProtectedItemOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ProtectedItemOperationResults.json
     */
    /**
     * Sample code: Get Operation Results of Protected Vm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationResultsOfProtectedVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectedItemOperationResults()
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
