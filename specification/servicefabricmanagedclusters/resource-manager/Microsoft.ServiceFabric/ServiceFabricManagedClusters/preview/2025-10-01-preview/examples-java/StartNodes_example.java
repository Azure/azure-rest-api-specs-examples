
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import java.util.Arrays;

/**
 * Samples for NodeTypes Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/StartNodes_example.json
     */
    /**
     * Sample code: Start nodes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void
        startNodes(com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().start("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withNodes(Arrays.asList("BE_0", "BE_1")), com.azure.core.util.Context.NONE);
    }
}
