
import com.azure.resourcemanager.networkcloud.models.KubernetesClusterRestartNodeParameters;

/**
 * Samples for KubernetesClusters RestartNode.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * KubernetesClusters_RestartNode.json
     */
    /**
     * Sample code: Restart a Kubernetes cluster node.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        restartAKubernetesClusterNode(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().restartNode("resourceGroupName", "kubernetesClusterName",
            new KubernetesClusterRestartNodeParameters().withNodeName("nodeName"), com.azure.core.util.Context.NONE);
    }
}
