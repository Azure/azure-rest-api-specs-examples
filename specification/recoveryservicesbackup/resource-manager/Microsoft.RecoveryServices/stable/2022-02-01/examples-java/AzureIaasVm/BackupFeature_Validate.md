```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureVMResourceFeatureSupportRequest;

/** Samples for FeatureSupport Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/BackupFeature_Validate.json
     */
    /**
     * Sample code: Check Azure Vm Backup Feature Support.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void checkAzureVmBackupFeatureSupport(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .featureSupports()
            .validateWithResponse(
                "southeastasia",
                new AzureVMResourceFeatureSupportRequest().withVmSize("Basic_A0").withVmSku("Premium"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.5/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
