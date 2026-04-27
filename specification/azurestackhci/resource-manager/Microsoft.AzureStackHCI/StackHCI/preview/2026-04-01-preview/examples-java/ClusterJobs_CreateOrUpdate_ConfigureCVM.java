
import com.azure.resourcemanager.azurestackhci.models.ConfidentialVmIntent;
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.HciConfigureCvmJobProperties;

/**
 * Samples for ClusterJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ClusterJobs_CreateOrUpdate_ConfigureCVM.json
     */
    /**
     * Sample code: ClusterJobs_CreateOrUpdate_ConfigureCVMJob.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        clusterJobsCreateOrUpdateConfigureCVMJob(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().define("ConfigureCVM").withExistingCluster("test-rg", "myCluster")
            .withProperties(new HciConfigureCvmJobProperties().withDeploymentMode(DeploymentMode.DEPLOY)
                .withConfidentialVmIntent(ConfidentialVmIntent.ENABLE))
            .create();
    }
}
