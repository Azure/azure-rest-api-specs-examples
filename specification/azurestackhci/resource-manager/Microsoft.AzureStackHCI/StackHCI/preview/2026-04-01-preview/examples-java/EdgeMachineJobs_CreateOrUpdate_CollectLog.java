
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.EdgeMachineJobProperties;

/**
 * Samples for EdgeMachineJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_CollectLog.json
     */
    /**
     * Sample code: EdgeMachineJobs_CreateOrUpdate_CollectLog.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsCreateOrUpdateCollectLog(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().define("triggerLogCollection").withExistingEdgeMachine("ArcInstance-rg", "machine1")
            .withProperties(new EdgeMachineJobProperties().withDeploymentMode(DeploymentMode.VALIDATE)).create();
    }
}
