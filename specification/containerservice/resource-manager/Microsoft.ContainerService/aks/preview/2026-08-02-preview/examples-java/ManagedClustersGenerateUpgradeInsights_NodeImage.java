
import com.azure.resourcemanager.containerservice.models.NodeImageGenerateUpgradeInsightsContent;

/**
 * Samples for UpgradeInsights Generate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_NodeImage.json
     */
    /**
     * Sample code: Generate a node OS image upgrade insight.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        generateANodeOSImageUpgradeInsight(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getUpgradeInsights()
            .generateWithResponse("rg1", "clustername1", new NodeImageGenerateUpgradeInsightsContent()
                .withAgentPoolName("nodepool1").withTargetNodeImageVersion("AKSUbuntu-2204gen2-202607.01.0"),
                com.azure.core.util.Context.NONE);
    }
}
