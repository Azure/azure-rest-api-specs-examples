
import com.azure.resourcemanager.hardwaresecuritymodules.models.RestoreRequestProperties;

/**
 * Samples for CloudHsmClusters ValidateRestoreProperties.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmCluster_RequestOrValidate_Restore_MaximumSet_Gen.json
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
