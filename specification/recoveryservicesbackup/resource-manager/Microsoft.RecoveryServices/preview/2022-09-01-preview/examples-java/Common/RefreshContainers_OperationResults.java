import com.azure.core.util.Context;

/** Samples for ProtectionContainerRefreshOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/RefreshContainers_OperationResults.json
     */
    /**
     * Sample code: Azure Vm Discovery Operation Result.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void azureVmDiscoveryOperationResult(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainerRefreshOperationResults()
            .getWithResponse(
                "NetSDKTestRsVault", "SwaggerTestRg", "Azure", "00000000-0000-0000-0000-000000000000", Context.NONE);
    }
}
