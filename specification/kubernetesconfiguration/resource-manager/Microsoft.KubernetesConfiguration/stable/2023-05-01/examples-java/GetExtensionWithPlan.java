/** Samples for Extensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/GetExtensionWithPlan.json
     */
    /**
     * Sample code: Get Extension with Plan.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void getExtensionWithPlan(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .getWithResponse(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "azureVote",
                com.azure.core.util.Context.NONE);
    }
}
