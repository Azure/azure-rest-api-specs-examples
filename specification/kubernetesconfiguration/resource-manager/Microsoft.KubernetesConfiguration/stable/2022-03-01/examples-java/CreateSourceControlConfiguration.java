import com.azure.core.util.Context;
import com.azure.resourcemanager.kubernetesconfiguration.fluent.models.SourceControlConfigurationInner;
import com.azure.resourcemanager.kubernetesconfiguration.models.HelmOperatorProperties;
import com.azure.resourcemanager.kubernetesconfiguration.models.OperatorScopeType;
import com.azure.resourcemanager.kubernetesconfiguration.models.OperatorType;
import java.util.HashMap;
import java.util.Map;

/** Samples for SourceControlConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateSourceControlConfiguration.json
     */
    /**
     * Sample code: Create Source Control Configuration.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void createSourceControlConfiguration(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .sourceControlConfigurations()
            .createOrUpdateWithResponse(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "SRS_GitHubConfig",
                new SourceControlConfigurationInner()
                    .withRepositoryUrl("git@github.com:k8sdeveloper425/flux-get-started")
                    .withOperatorNamespace("SRS_Namespace")
                    .withOperatorInstanceName("SRSGitHubFluxOp-01")
                    .withOperatorType(OperatorType.FLUX)
                    .withOperatorParams("--git-email=xyzgituser@users.srs.github.com")
                    .withConfigurationProtectedSettings(mapOf("protectedSetting1Key", "protectedSetting1Value"))
                    .withOperatorScope(OperatorScopeType.NAMESPACE)
                    .withSshKnownHostsContents(
                        "c3NoLmRldi5henVyZS5jb20gc3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFDN0hyMW9UV3FOcU9sekdKT2ZHSjROYWtWeUl6ZjFyWFlkNGQ3d282akJsa0x2Q0E0b2RCbEwwbURVeVowL1FVZlRUcWV1K3RtMjJnT3N2K1ZyVlRNazZ2d1JVNzVnWS95OXV0NU1iM2JSNUJWNThkS1h5cTlBOVVlQjVDYWtlaG41WmdtNngxbUtvVnlmK0ZGbjI2aVlxWEpSZ3pJWlpjWjVWNmhyRTBRZzM5a1ptNGF6NDhvMEFVYmY2U3A0U0xkdm51TWEyc1ZOd0hCYm9TN0VKa201N1hRUFZVMy9RcHlOTEhiV0Rkend0cmxTK2V6MzBTM0FkWWhMS0VPeEFHOHdlT255cnRMSkFVZW45bVRrb2w4b0lJMWVkZjdtV1diV1ZmMG5CbWx5MjErblpjbUNUSVNRQnRkY3lQYUVubzdmRlFNREQyNi9zMGxmS29iNEt3OEg=")
                    .withEnableHelmOperator(true)
                    .withHelmOperatorProperties(
                        new HelmOperatorProperties()
                            .withChartVersion("0.3.0")
                            .withChartValues(
                                "--set git.ssh.secretName=flux-git-deploy --set tillerNamespace=kube-system")),
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
