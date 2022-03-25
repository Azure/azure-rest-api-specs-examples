Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kubernetesconfiguration_1.0.0-beta.3/sdk/kubernetesconfiguration/azure-resourcemanager-kubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kubernetesconfiguration.models.FluxConfigurationPatch;
import com.azure.resourcemanager.kubernetesconfiguration.models.GitRepositoryPatchDefinition;
import com.azure.resourcemanager.kubernetesconfiguration.models.KustomizationPatchDefinition;
import java.util.HashMap;
import java.util.Map;

/** Samples for FluxConfigurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/PatchFluxConfiguration.json
     */
    /**
     * Sample code: Patch Flux Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void patchFluxConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigurations()
            .update(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "srs-fluxconfig",
                new FluxConfigurationPatch()
                    .withSuspend(true)
                    .withGitRepository(
                        new GitRepositoryPatchDefinition()
                            .withUrl("https://github.com/jonathan-innis/flux2-kustomize-helm-example.git"))
                    .withKustomizations(
                        mapOf(
                            "srs-kustomization1",
                            null,
                            "srs-kustomization2",
                            new KustomizationPatchDefinition()
                                .withPath("./test/alt-path")
                                .withSyncIntervalInSeconds(300L),
                            "srs-kustomization3",
                            new KustomizationPatchDefinition()
                                .withPath("./test/another-path")
                                .withSyncIntervalInSeconds(300L))),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
