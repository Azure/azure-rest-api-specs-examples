import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureVMResourceFeatureSupportRequest;

/** Samples for FeatureSupport Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/BackupFeature_Validate.json
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
