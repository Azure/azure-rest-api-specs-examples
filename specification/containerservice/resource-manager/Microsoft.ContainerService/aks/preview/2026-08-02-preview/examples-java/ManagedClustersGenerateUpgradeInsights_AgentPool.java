
import com.azure.resourcemanager.containerservice.models.AgentPoolGenerateUpgradeInsightsContent;

/**
 * Samples for UpgradeInsights Generate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_AgentPool.json
     */
    /**
     * Sample code: Generate a single agent pool Kubernetes version upgrade insight.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void generateASingleAgentPoolKubernetesVersionUpgradeInsight(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getUpgradeInsights()
            .generateWithResponse("rg1", "clustername1", new AgentPoolGenerateUpgradeInsightsContent()
                .withAgentPoolName("nodepool1").withTargetKubernetesVersion("1.33.5"),
                com.azure.core.util.Context.NONE);
    }
}
