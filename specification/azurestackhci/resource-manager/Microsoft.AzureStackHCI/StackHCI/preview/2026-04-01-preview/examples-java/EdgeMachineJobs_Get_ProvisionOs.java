
/**
 * Samples for EdgeMachineJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_Get_ProvisionOs.json
     */
    /**
     * Sample code: EdgeMachineJobs_Get_ProvisionOs.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsGetProvisionOs(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().getWithResponse("ArcInstance-rg", "machine1", "ProvisionOs",
            com.azure.core.util.Context.NONE);
    }
}
