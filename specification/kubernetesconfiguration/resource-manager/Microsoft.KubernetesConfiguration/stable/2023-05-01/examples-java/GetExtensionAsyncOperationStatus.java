/** Samples for OperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/GetExtensionAsyncOperationStatus.json
     */
    /**
     * Sample code: ExtensionAsyncOperationStatus Get.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void extensionAsyncOperationStatusGet(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .operationStatus()
            .getWithResponse(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "ClusterMonitor",
                "99999999-9999-9999-9999-999999999999",
                com.azure.core.util.Context.NONE);
    }
}
