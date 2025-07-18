
import com.azure.resourcemanager.hardwaresecuritymodules.models.RestoreRequestProperties;

/**
 * Samples for CloudHsmClusters ValidateRestoreProperties.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-03-31/CloudHsmCluster_RequestOrValidate_Restore_MaximumSet_Gen_ValidateRestoreProperties.json
     */
    /**
     * Sample code: CloudHsmCluster_ValidateRestore_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterValidateRestoreMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusters().validateRestoreProperties("rgcloudhsm", "chsm1",
            new RestoreRequestProperties()
                .withAzureStorageBlobContainerUri("https://myaccount.blob.core.windows.net/sascontainer/sasContainer")
                .withBackupId("backupId"),
            com.azure.core.util.Context.NONE);
    }
}
