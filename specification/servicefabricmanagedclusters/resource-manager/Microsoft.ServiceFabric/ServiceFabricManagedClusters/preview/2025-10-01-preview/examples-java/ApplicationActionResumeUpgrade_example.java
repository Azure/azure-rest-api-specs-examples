
import com.azure.resourcemanager.servicefabricmanagedclusters.models.RuntimeResumeApplicationUpgradeParameters;

/**
 * Samples for Applications ResumeUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ApplicationActionResumeUpgrade_example.json
     */
    /**
     * Sample code: Resume upgrade.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void resumeUpgrade(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().resumeUpgrade("resRg", "myCluster", "myApp",
            new RuntimeResumeApplicationUpgradeParameters().withUpgradeDomainName("UD1"),
            com.azure.core.util.Context.NONE);
    }
}
