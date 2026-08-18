
import com.azure.resourcemanager.recoveryservicesbackup.models.ProtectedItemConfigureSourceScanRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.SourceScanAction;

/**
 * Samples for ConfigureSourceScan Execute.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AzureIaasVm/ConfigureSourceScan.json
     */
    /**
     * Sample code: Configure Source Scan for Protected Item.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void configureSourceScanForProtectedItem(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.configureSourceScans().execute("SwaggerTestRg", "NetSDKTestRsVault", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            new ProtectedItemConfigureSourceScanRequest().withSourceScanAction(SourceScanAction.ENABLE),
            com.azure.core.util.Context.NONE);
    }
}
