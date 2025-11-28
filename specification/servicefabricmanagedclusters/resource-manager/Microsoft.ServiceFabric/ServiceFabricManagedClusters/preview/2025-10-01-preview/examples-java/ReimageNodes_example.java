
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import java.util.Arrays;

/**
 * Samples for NodeTypes Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ReimageNodes_example.json
     */
    /**
     * Sample code: Reimage nodes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void reimageNodes(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().reimage("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withNodes(Arrays.asList("BE_0", "BE_3")), com.azure.core.util.Context.NONE);
    }
}
