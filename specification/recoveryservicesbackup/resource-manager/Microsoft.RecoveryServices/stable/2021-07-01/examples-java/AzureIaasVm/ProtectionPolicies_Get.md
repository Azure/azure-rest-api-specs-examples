```java
import com.azure.core.util.Context;

/** Samples for ProtectionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/AzureIaasVm/ProtectionPolicies_Get.json
     */
    /**
     * Sample code: Get Azure IaasVm Protection Policy Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureIaasVmProtectionPolicyDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().getWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "testPolicy1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
