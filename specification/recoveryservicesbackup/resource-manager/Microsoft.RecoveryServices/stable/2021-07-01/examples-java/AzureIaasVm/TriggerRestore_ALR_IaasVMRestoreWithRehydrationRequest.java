import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;

/** Samples for Restores Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreWithRehydrationRequest.json
     */
    /**
     * Sample code: Restore to New Azure IaasVm with IaasVMRestoreWithRehydrationRequest.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreToNewAzureIaasVmWithIaasVMRestoreWithRehydrationRequest(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .restores()
            .trigger(
                "testVault",
                "netsdktestrg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
                "348916168024334",
                new RestoreRequestResource().withProperties(new RestoreRequest()),
                Context.NONE);
    }
}
