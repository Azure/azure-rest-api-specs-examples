
import com.azure.resourcemanager.hardwaresecuritymodules.models.BackupRequestProperties;

/**
 * Samples for CloudHsmClusters ValidateBackupProperties.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmCluster_CreateOrValidate_Backup_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_ValidateBackup_Validation_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterValidateBackupValidationMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusters().validateBackupProperties("rgcloudhsm", "chsm1",
            new BackupRequestProperties()
                .withAzureStorageBlobContainerUri("https://myaccount.blob.core.windows.net/sascontainer/sasContainer")
                .withToken("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
