import com.azure.resourcemanager.kubernetesconfiguration.fluent.models.FluxConfigurationInner;
import com.azure.resourcemanager.kubernetesconfiguration.models.BucketDefinition;
import com.azure.resourcemanager.kubernetesconfiguration.models.KustomizationDefinition;
import com.azure.resourcemanager.kubernetesconfiguration.models.ScopeType;
import com.azure.resourcemanager.kubernetesconfiguration.models.SourceKindType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for FluxConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateFluxConfigurationWithBucket.json
     */
    /**
     * Sample code: Create Flux Configuration with Bucket Source Kind.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void createFluxConfigurationWithBucketSourceKind(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .fluxConfigurations()
            .createOrUpdate(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "srs-fluxconfig",
                new FluxConfigurationInner()
                    .withScope(ScopeType.CLUSTER)
                    .withNamespace("srs-namespace")
                    .withSourceKind(SourceKindType.BUCKET)
                    .withSuspend(false)
                    .withBucket(
                        new BucketDefinition()
                            .withUrl("https://fluxminiotest.az.minio.io")
                            .withBucketName("flux")
                            .withTimeoutInSeconds(1000L)
                            .withSyncIntervalInSeconds(1000L)
                            .withAccessKey("fakeTokenPlaceholder"))
                    .withKustomizations(
                        mapOf(
                            "srs-kustomization1",
                            new KustomizationDefinition()
                                .withPath("./test/path")
                                .withDependsOn(Arrays.asList())
                                .withTimeoutInSeconds(600L)
                                .withSyncIntervalInSeconds(600L),
                            "srs-kustomization2",
                            new KustomizationDefinition()
                                .withPath("./other/test/path")
                                .withDependsOn(Arrays.asList("srs-kustomization1"))
                                .withTimeoutInSeconds(600L)
                                .withSyncIntervalInSeconds(600L)
                                .withRetryIntervalInSeconds(600L)
                                .withPrune(false))),
                com.azure.core.util.Context.NONE);
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
