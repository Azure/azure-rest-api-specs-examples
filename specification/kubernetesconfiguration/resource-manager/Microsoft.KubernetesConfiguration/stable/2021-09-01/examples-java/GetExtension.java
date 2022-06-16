import com.azure.core.util.Context;
import com.azure.resourcemanager.kubernetesconfiguration.models.ExtensionsClusterResourceName;
import com.azure.resourcemanager.kubernetesconfiguration.models.ExtensionsClusterRp;

/** Samples for Extensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2021-09-01/examples/GetExtension.json
     */
    /**
     * Sample code: Get Extension.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void getExtension(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .getWithResponse(
                "rg1",
                ExtensionsClusterRp.MICROSOFT_KUBERNETES,
                ExtensionsClusterResourceName.CONNECTED_CLUSTERS,
                "clusterName1",
                "ClusterMonitor",
                Context.NONE);
    }
}
