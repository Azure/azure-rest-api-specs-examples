/** Samples for JobCancellations Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/TriggerCancelJob.json
     */
    /**
     * Sample code: Cancel Job.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void cancelJob(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .jobCancellations()
            .triggerWithResponse(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "00000000-0000-0000-0000-000000000000",
                com.azure.core.util.Context.NONE);
    }
}
