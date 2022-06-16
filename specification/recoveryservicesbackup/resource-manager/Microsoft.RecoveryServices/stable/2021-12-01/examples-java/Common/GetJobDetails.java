import com.azure.core.util.Context;

/** Samples for JobDetails Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/GetJobDetails.json
     */
    /**
     * Sample code: Get Job Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getJobDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .jobDetails()
            .getWithResponse(
                "NetSDKTestRsVault", "SwaggerTestRg", "00000000-0000-0000-0000-000000000000", Context.NONE);
    }
}
