/** Samples for OperationStatus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/ListAsyncOperationStatus.json
     */
    /**
     * Sample code: AsyncOperationStatus List.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void asyncOperationStatusList(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .operationStatus()
            .list("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", com.azure.core.util.Context.NONE);
    }
}
