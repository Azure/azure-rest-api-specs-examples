
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVmilrRegistrationRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.IlrRequestResource;

/**
 * Samples for ItemLevelRecoveryConnections Provision.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * AzureIaasVm/Provision_Ilr.json
     */
    /**
     * Sample code: Provision Instant Item Level Recovery for Azure Vm.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void provisionInstantItemLevelRecoveryForAzureVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.itemLevelRecoveryConnections().provisionWithResponse("PySDKBackupTestRsVault", "PythonSDKBackupTestRg",
            "Azure", "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
            "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", "1",
            new IlrRequestResource().withProperties(new IaasVmilrRegistrationRequest()
                .withRecoveryPointId("38823086363464")
                .withVirtualMachineId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/pysdktestrg/providers/Microsoft.Compute/virtualMachines/pysdktestv2vm1")
                .withInitiatorName("Hello World").withRenewExistingRegistration(true)),
            com.azure.core.util.Context.NONE);
    }
}
