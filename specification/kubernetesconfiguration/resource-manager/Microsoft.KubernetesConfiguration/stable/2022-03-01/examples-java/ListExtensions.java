import com.azure.core.util.Context;

/** Samples for Extensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/ListExtensions.json
     */
    /**
     * Sample code: List Extensions.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void listExtensions(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager.extensions().list("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", Context.NONE);
    }
}
