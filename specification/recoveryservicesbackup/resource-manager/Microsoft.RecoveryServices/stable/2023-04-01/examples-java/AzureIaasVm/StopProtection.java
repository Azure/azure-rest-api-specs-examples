import com.azure.resourcemanager.recoveryservicesbackup.models.AzureIaaSComputeVMProtectedItem;
import com.azure.resourcemanager.recoveryservicesbackup.models.ProtectionState;

/** Samples for ProtectedItems CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/StopProtection.json
     */
    /**
     * Sample code: Stop Protection with retain data on Azure IaasVm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void stopProtectionWithRetainDataOnAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectedItems()
            .define("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1")
            .withRegion((String) null)
            .withExistingProtectionContainer(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "Azure",
                "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1")
            .withProperties(
                new AzureIaaSComputeVMProtectedItem()
                    .withSourceResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                    .withProtectionState(ProtectionState.PROTECTION_STOPPED))
            .create();
    }
}
