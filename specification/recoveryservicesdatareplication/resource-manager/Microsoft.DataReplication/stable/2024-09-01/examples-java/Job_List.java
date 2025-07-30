
/**
 * Samples for Job List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Job_List.json
     */
    /**
     * Sample code: Lists the jobs.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheJobs(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.jobs().list("rgrecoveryservicesdatareplication", "4", null, "rdavrzbethhslmkqgajontnxsue", null,
            com.azure.core.util.Context.NONE);
    }
}
