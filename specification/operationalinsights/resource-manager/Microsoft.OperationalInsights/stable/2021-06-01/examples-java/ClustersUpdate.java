import com.azure.resourcemanager.loganalytics.models.Capacity;
import com.azure.resourcemanager.loganalytics.models.Cluster;
import com.azure.resourcemanager.loganalytics.models.ClusterSku;
import com.azure.resourcemanager.loganalytics.models.ClusterSkuNameEnum;
import com.azure.resourcemanager.loganalytics.models.Identity;
import com.azure.resourcemanager.loganalytics.models.IdentityType;
import com.azure.resourcemanager.loganalytics.models.KeyVaultProperties;
import com.azure.resourcemanager.loganalytics.models.UserIdentityProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersUpdate.json
     */
    /**
     * Sample code: ClustersPatch.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void clustersPatch(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("oiautorest6685", "oiautorest6685", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "val1"))
            .withIdentity(
                new Identity()
                    .withType(IdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
                            new UserIdentityProperties())))
            .withSku(
                new ClusterSku()
                    .withCapacity(Capacity.ONE_ZERO_ZERO_ZERO)
                    .withName(ClusterSkuNameEnum.CAPACITY_RESERVATION))
            .withKeyVaultProperties(
                new KeyVaultProperties()
                    .withKeyVaultUri("fakeTokenPlaceholder")
                    .withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder")
                    .withKeyRsaSize(1024))
            .apply();
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
