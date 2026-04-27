
/**
 * Samples for EdgeMachineJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: EdgeMachineJobs_List_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsListMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().list("ArcInstance-rg", "machine1", com.azure.core.util.Context.NONE);
    }
}
