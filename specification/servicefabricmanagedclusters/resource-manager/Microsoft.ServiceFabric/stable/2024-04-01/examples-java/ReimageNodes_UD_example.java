
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.UpdateType;

/**
 * Samples for NodeTypes Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/
     * ReimageNodes_UD_example.json
     */
    /**
     * Sample code: Reimage all nodes By upgrade domain.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void reimageAllNodesByUpgradeDomain(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().reimage("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withUpdateType(UpdateType.BY_UPGRADE_DOMAIN),
            com.azure.core.util.Context.NONE);
    }
}
