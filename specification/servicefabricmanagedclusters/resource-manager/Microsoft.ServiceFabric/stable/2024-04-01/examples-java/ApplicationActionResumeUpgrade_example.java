
import com.azure.resourcemanager.servicefabricmanagedclusters.models.RuntimeResumeApplicationUpgradeParameters;

/**
 * Samples for Applications ResumeUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/
     * ApplicationActionResumeUpgrade_example.json
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
