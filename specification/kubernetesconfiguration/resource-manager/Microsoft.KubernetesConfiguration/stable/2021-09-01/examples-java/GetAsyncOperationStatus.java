import com.azure.core.util.Context;
import com.azure.resourcemanager.kubernetesconfiguration.models.ExtensionsClusterResourceName;
import com.azure.resourcemanager.kubernetesconfiguration.models.ExtensionsClusterRp;

/** Samples for OperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2021-09-01/examples/GetAsyncOperationStatus.json
     */
    /**
     * Sample code: AsyncOperationStatus Get.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void asyncOperationStatusGet(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .operationStatus()
            .getWithResponse(
                "rg1",
                ExtensionsClusterRp.MICROSOFT_KUBERNETES,
                ExtensionsClusterResourceName.CONNECTED_CLUSTERS,
                "clusterName1",
                "ClusterMonitor",
                "99999999-9999-9999-9999-999999999999",
                Context.NONE);
    }
}
