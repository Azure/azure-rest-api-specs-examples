
/**
 * Samples for EdgeMachineJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineJobs_Get_CollectLog.json
     */
    /**
     * Sample code: EdgeMachineJobs_Get_CollectLog.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsGetCollectLog(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().getWithResponse("ArcInstance-rg", "machine1", "collectLog",
            com.azure.core.util.Context.NONE);
    }
}
