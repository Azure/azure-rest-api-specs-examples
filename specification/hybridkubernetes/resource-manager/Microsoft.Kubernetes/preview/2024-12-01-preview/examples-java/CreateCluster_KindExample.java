
import com.azure.resourcemanager.hybridkubernetes.models.AadProfile;
import com.azure.resourcemanager.hybridkubernetes.models.ArcAgentProfile;
import com.azure.resourcemanager.hybridkubernetes.models.AutoUpgradeOptions;
import com.azure.resourcemanager.hybridkubernetes.models.AzureHybridBenefit;
import com.azure.resourcemanager.hybridkubernetes.models.ConnectedClusterIdentity;
import com.azure.resourcemanager.hybridkubernetes.models.ConnectedClusterKind;
import com.azure.resourcemanager.hybridkubernetes.models.OidcIssuerProfile;
import com.azure.resourcemanager.hybridkubernetes.models.ResourceIdentityType;
import com.azure.resourcemanager.hybridkubernetes.models.SystemComponent;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConnectedCluster CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * CreateCluster_KindExample.json
     */
    /**
     * Sample code: CreateCluster_KindExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void
        createClusterKindExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().define("testCluster").withRegion("East US").withExistingResourceGroup("k8sc-rg")
            .withIdentity(new ConnectedClusterIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withAgentPublicKeyCertificate("").withTags(mapOf()).withKind(ConnectedClusterKind.PROVISIONED_CLUSTER)
            .withDistribution("AKS").withDistributionVersion("1.0")
            .withAzureHybridBenefit(AzureHybridBenefit.NOT_APPLICABLE)
            .withAadProfile(new AadProfile().withEnableAzureRbac(true)
                .withAdminGroupObjectIDs(Arrays.asList("56f988bf-86f1-41af-91ab-2d7cd011db47"))
                .withTenantId("82f988bf-86f1-41af-91ab-2d7cd011db47"))
            .withArcAgentProfile(new ArcAgentProfile().withDesiredAgentVersion("0.1.0")
                .withAgentAutoUpgrade(AutoUpgradeOptions.ENABLED)
                .withSystemComponents(Arrays.asList(
                    new SystemComponent().withType("Strato").withUserSpecifiedVersion("0.1.1").withMajorVersion(0))))
            .withOidcIssuerProfile(new OidcIssuerProfile().withEnabled(true)).create();
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
