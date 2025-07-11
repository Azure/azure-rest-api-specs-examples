
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import java.util.Arrays;

/**
 * Samples for NodeTypes DeleteNode.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/DeleteNodes_example.json
     */
    /**
     * Sample code: Delete nodes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteNodes(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().deleteNode("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withNodes(Arrays.asList("BE_0", "BE_3")), com.azure.core.util.Context.NONE);
    }
}
