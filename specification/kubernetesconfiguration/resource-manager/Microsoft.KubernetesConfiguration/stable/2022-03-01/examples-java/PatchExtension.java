import com.azure.core.util.Context;
import com.azure.resourcemanager.kubernetesconfiguration.models.PatchExtension;
import java.util.HashMap;
import java.util.Map;

/** Samples for Extensions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/PatchExtension.json
     */
    /**
     * Sample code: Update Extension.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void updateExtension(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .update(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "ClusterMonitor",
                new PatchExtension()
                    .withAutoUpgradeMinorVersion(true)
                    .withReleaseTrain("Preview")
                    .withConfigurationSettings(
                        mapOf(
                            "omsagent.env.clusterName",
                            "clusterName1",
                            "omsagent.secret.wsid",
                            "a38cef99-5a89-52ed-b6db-22095c23664b"))
                    .withConfigurationProtectedSettings(mapOf("omsagent.secret.key", "secretKeyValue01")),
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
