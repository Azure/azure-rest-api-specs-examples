import com.azure.core.util.Context;
import com.azure.resourcemanager.kubernetesconfiguration.models.ExtensionsClusterResourceName;
import com.azure.resourcemanager.kubernetesconfiguration.models.ExtensionsClusterRp;

/** Samples for Extensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2021-09-01/examples/ListExtensions.json
     */
    /**
     * Sample code: List Extensions.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void listExtensions(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .list(
                "rg1",
                ExtensionsClusterRp.MICROSOFT_KUBERNETES,
                ExtensionsClusterResourceName.CONNECTED_CLUSTERS,
                "clusterName1",
                Context.NONE);
    }
}
