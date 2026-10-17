
import com.azure.resourcemanager.azurestackhci.models.AddServerJobServerDetails;
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.EceDeploymentSecrets;
import com.azure.resourcemanager.azurestackhci.models.EceSecrets;
import com.azure.resourcemanager.azurestackhci.models.HciAddServerJobProperties;
import java.util.Arrays;

/**
 * Samples for ClusterJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ClusterJobs_CreateOrUpdate_AddServer.json
     */
    /**
     * Sample code: ClusterJobs_CreateOrUpdate_AddServerJob.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        clusterJobsCreateOrUpdateAddServerJob(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().define("addserver").withExistingCluster("test-rg", "myCluster")
            .withProperties(new HciAddServerJobProperties().withDeploymentMode(DeploymentMode.DEPLOY)
                .withAddServerJobServerDetails(Arrays.asList(new AddServerJobServerDetails().withServerName("Node-2")
                    .withHostIpv4Address("192.168.10.10").withServerResourceId(
                        "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2")))
                .withSecrets(Arrays.asList(new EceDeploymentSecrets().withSecretName("fakeTokenPlaceholder")
                    .withEceSecretName(EceSecrets.LOCAL_ADMIN_CREDENTIAL).withSecretLocation("fakeTokenPlaceholder"),
                    new EceDeploymentSecrets().withSecretName("fakeTokenPlaceholder")
                        .withEceSecretName(EceSecrets.AZURE_STACK_LCMUSER_CREDENTIAL)
                        .withSecretLocation("fakeTokenPlaceholder"))))
            .create();
    }
}
