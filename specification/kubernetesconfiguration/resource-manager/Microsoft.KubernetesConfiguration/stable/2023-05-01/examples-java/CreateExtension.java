import com.azure.resourcemanager.kubernetesconfiguration.fluent.models.ExtensionInner;
import com.azure.resourcemanager.kubernetesconfiguration.models.Scope;
import com.azure.resourcemanager.kubernetesconfiguration.models.ScopeCluster;
import java.util.HashMap;
import java.util.Map;

/** Samples for Extensions Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/CreateExtension.json
     */
    /**
     * Sample code: Create Extension.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void createExtension(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .create(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "ClusterMonitor",
                new ExtensionInner()
                    .withExtensionType("azuremonitor-containers")
                    .withAutoUpgradeMinorVersion(true)
                    .withReleaseTrain("Preview")
                    .withScope(new Scope().withCluster(new ScopeCluster().withReleaseNamespace("kube-system")))
                    .withConfigurationSettings(
                        mapOf(
                            "omsagent.env.clusterName", "clusterName1", "omsagent.secret.wsid", "fakeTokenPlaceholder"))
                    .withConfigurationProtectedSettings(mapOf("omsagent.secret.key", "fakeTokenPlaceholder")),
                com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
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
