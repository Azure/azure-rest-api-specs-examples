
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeTypeActionParameters;
import java.util.Arrays;

/**
 * Samples for NodeTypes Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/RestartNodes_example.json
     */
    /**
     * Sample code: Restart nodes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void restartNodes(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().restart("resRg", "myCluster", "BE",
            new NodeTypeActionParameters().withNodes(Arrays.asList("BE_0", "BE_3")), com.azure.core.util.Context.NONE);
    }
}
