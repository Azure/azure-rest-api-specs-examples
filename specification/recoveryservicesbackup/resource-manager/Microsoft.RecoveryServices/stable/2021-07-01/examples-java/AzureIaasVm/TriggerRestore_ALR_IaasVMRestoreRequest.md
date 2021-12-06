Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;

/** Samples for Restores Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreRequest.json
     */
    /**
     * Sample code: Restore to New Azure IaasVm with IaasVMRestoreRequest.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreToNewAzureIaasVmWithIaasVMRestoreRequest(
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
                new RestoreRequestResource()
                    .withProperties(
                        new IaasVMRestoreRequest()
                            .withRecoveryPointId("348916168024334")
                            .withRecoveryType(RecoveryType.ALTERNATE_LOCATION)
                            .withSourceResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                            .withTargetVirtualMachineId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2/providers/Microsoft.Compute/virtualmachines/RSMDALRVM981435")
                            .withTargetResourceGroupId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2")
                            .withStorageAccountId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount")
                            .withVirtualNetworkId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet")
                            .withSubnetId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet/subnets/default")
                            .withRegion("southeastasia")
                            .withCreateNewCloudService(false)
                            .withOriginalStorageAccountOption(false)
                            .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                            .withIdentityInfo(new IdentityInfo().withIsSystemAssignedIdentity(true))),
                Context.NONE);
    }
}
```
