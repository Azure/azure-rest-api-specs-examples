Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.recoveryservicesbackup.models.ProtectedItem;

/** Samples for ProtectedItems CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureIaasVm/StopProtection.json
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
                new ProtectedItem()
                    .withSourceResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"))
            .create();
    }
}
```
