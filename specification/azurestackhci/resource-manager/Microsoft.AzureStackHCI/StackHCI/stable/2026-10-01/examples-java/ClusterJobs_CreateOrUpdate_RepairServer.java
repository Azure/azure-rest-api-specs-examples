
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.EceDeploymentSecrets;
import com.azure.resourcemanager.azurestackhci.models.EceSecrets;
import com.azure.resourcemanager.azurestackhci.models.HciRepairServerJobProperties;
import com.azure.resourcemanager.azurestackhci.models.RepairServerJobServerDetails;
import java.util.Arrays;

/**
 * Samples for ClusterJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ClusterJobs_CreateOrUpdate_RepairServer.json
     */
    /**
     * Sample code: ClusterJobs_CreateOrUpdate_RepairServerJob.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        clusterJobsCreateOrUpdateRepairServerJob(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().define("repairserver").withExistingCluster("test-rg", "myCluster")
            .withProperties(new HciRepairServerJobProperties().withDeploymentMode(DeploymentMode.DEPLOY)
                .withRepairServerJobServerDetails(
                    Arrays.asList(new RepairServerJobServerDetails().withServerName("Node-2").withServerResourceId(
                        "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2")))
                .withSecrets(Arrays.asList(new EceDeploymentSecrets().withSecretName("fakeTokenPlaceholder")
                    .withEceSecretName(EceSecrets.LOCAL_ADMIN_CREDENTIAL).withSecretLocation("fakeTokenPlaceholder"),
                    new EceDeploymentSecrets().withSecretName("fakeTokenPlaceholder")
                        .withEceSecretName(EceSecrets.AZURE_STACK_LCMUSER_CREDENTIAL)
                        .withSecretLocation("fakeTokenPlaceholder"))))
            .create();
    }
}
