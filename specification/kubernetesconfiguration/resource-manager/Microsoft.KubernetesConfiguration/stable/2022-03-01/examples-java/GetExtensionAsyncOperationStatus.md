```java
import com.azure.core.util.Context;

/** Samples for OperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/GetExtensionAsyncOperationStatus.json
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
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kubernetesconfiguration_1.0.0-beta.3/sdk/kubernetesconfiguration/azure-resourcemanager-kubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.
