import com.azure.resourcemanager.kubernetesconfiguration.fluent.models.ExtensionInner;
import com.azure.resourcemanager.kubernetesconfiguration.models.Plan;
import java.util.HashMap;
import java.util.Map;

/** Samples for Extensions Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateExtensionWithPlan.json
     */
    /**
     * Sample code: Create Extension with Plan.
     *
     * @param manager Entry point to SourceControlConfigurationManager.
     */
    public static void createExtensionWithPlan(
        com.azure.resourcemanager.kubernetesconfiguration.SourceControlConfigurationManager manager) {
        manager
            .extensions()
            .create(
                "rg1",
                "Microsoft.Kubernetes",
                "connectedClusters",
                "clusterName1",
                "azureVote",
                new ExtensionInner()
                    .withPlan(
                        new Plan()
                            .withName("azure-vote-standard")
                            .withPublisher("Microsoft")
                            .withProduct("azure-vote-standard-offer-id"))
                    .withExtensionType("azure-vote")
                    .withAutoUpgradeMinorVersion(true)
                    .withReleaseTrain("Preview"),
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
