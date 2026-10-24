
import com.azure.resourcemanager.containerservice.models.ControlPlaneAndAgentPoolsGenerateUpgradeInsightsContent;

/**
 * Samples for UpgradeInsights Generate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_ControlPlaneAndAgentPools.json
     */
    /**
     * Sample code: Generate a combined control plane and agent pools upgrade insight.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void generateACombinedControlPlaneAndAgentPoolsUpgradeInsight(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getUpgradeInsights().generateWithResponse("rg1", "clustername1",
            new ControlPlaneAndAgentPoolsGenerateUpgradeInsightsContent().withTargetKubernetesVersion("1.33.5"),
            com.azure.core.util.Context.NONE);
    }
}
