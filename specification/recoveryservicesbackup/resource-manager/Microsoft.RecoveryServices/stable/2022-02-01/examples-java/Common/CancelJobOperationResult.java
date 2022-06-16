import com.azure.core.util.Context;

/** Samples for JobOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/Common/CancelJobOperationResult.json
     */
    /**
     * Sample code: Cancel Job Operation Result.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void cancelJobOperationResult(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .jobOperationResults()
            .getWithResponse(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "00000000-0000-0000-0000-000000000000",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
