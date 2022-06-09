```java
import com.azure.core.util.Context;

/** Samples for FluxConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/DeleteFluxConfiguration.json
     */
    /**
     * Sample code: Delete Flux Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void deleteFluxConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigurations()
            .delete(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "srs-fluxconfig",
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kubernetesconfiguration_1.0.0-beta.3/sdk/kubernetesconfiguration/azure-resourcemanager-kubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.
