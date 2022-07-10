import com.azure.resourcemanager.loganalytics.models.Capacity;
import com.azure.resourcemanager.loganalytics.models.ClusterSku;
import com.azure.resourcemanager.loganalytics.models.ClusterSkuNameEnum;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersCreate.json
     */
    /**
     * Sample code: ClustersCreate.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void clustersCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .clusters()
            .define("oiautorest6685")
            .withRegion("australiasoutheast")
            .withExistingResourceGroup("oiautorest6685")
            .withTags(mapOf("tag1", "val1"))
            .withSku(
                new ClusterSku()
                    .withCapacity(Capacity.ONE_ZERO_ZERO_ZERO)
                    .withName(ClusterSkuNameEnum.CAPACITY_RESERVATION))
            .create();
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
