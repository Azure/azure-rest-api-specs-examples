
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.UpdateType;

/**
 * Samples for NodeTypes Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/RedeployNodes_UD_example.json
     */
    /**
     * Sample code: Redeploy all nodes by upgrade domain.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void redeployAllNodesByUpgradeDomain(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().redeploy("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withUpdateType(UpdateType.BY_UPGRADE_DOMAIN),
            com.azure.core.util.Context.NONE);
    }
}
