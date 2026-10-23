
import com.azure.resourcemanager.containerservice.models.ControlPlaneGenerateUpgradeInsightsContent;

/**
 * Samples for UpgradeInsights Generate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_ControlPlane.json
     */
    /**
     * Sample code: Generate a control plane upgrade insight.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void generateAControlPlaneUpgradeInsight(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getUpgradeInsights().generateWithResponse("rg1", "clustername1",
            new ControlPlaneGenerateUpgradeInsightsContent().withTargetKubernetesVersion("1.33.5"),
            com.azure.core.util.Context.NONE);
    }
}
