
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import java.util.Arrays;

/**
 * Samples for NodeTypes Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/RedeployNodes_example.json
     */
    /**
     * Sample code: Redeploy nodes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void redeployNodes(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().redeploy("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withNodes(Arrays.asList("BE_0", "BE_1")), com.azure.core.util.Context.NONE);
    }
}
