
/**
 * Samples for EdgeMachineJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: EdgeMachineJobs_Delete_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsDeleteMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().delete("ArcInstance-rg", "machine1", "triggerLogCollection",
            com.azure.core.util.Context.NONE);
    }
}
