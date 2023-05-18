/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/OperationsList.json
     */
    /**
     * Sample code: BatchAccountDelete.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void batchAccountDelete(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
