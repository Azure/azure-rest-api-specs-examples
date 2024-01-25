
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureIaaSComputeVMProtectedItem;

/**
 * Samples for ProtectedItems CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * AzureIaasVm/ConfigureProtection.json
     */
    /**
     * Sample code: Enable Protection on Azure IaasVm.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void enableProtectionOnAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectedItems().define("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1").withRegion((String) null)
            .withExistingProtectionContainer("NetSDKTestRsVault", "SwaggerTestRg", "Azure",
                "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1")
            .withProperties(new AzureIaaSComputeVMProtectedItem().withSourceResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                .withPolicyId(
                    "/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/DefaultPolicy"))
            .create();
    }
}
