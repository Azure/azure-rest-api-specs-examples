
import com.azure.resourcemanager.cloudhealth.models.EvaluationRule;
import com.azure.resourcemanager.cloudhealth.models.MetricAggregationType;
import com.azure.resourcemanager.cloudhealth.models.RefreshInterval;
import com.azure.resourcemanager.cloudhealth.models.ResourceMetricSignalDefinitionProperties;
import com.azure.resourcemanager.cloudhealth.models.SignalOperator;
import com.azure.resourcemanager.cloudhealth.models.ThresholdRule;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SignalDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/SignalDefinitions_CreateOrUpdate.json
     */
    /**
     * Sample code: SignalDefinitions_CreateOrUpdate.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        signalDefinitionsCreateOrUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.signalDefinitions().define("sig1").withExistingHealthmodel("rgopenapi", "myHealthModel")
            .withProperties(new ResourceMetricSignalDefinitionProperties().withDisplayName("cpu usage")
                .withRefreshInterval(RefreshInterval.PT1M).withLabels(mapOf("key4788", "fakeTokenPlaceholder"))
                .withDataUnit("byte")
                .withEvaluationRules(new EvaluationRule()
                    .withDegradedRule(new ThresholdRule().withOperator(SignalOperator.LOWER_THAN).withThreshold("65"))
                    .withUnhealthyRule(new ThresholdRule().withOperator(SignalOperator.LOWER_THAN).withThreshold("60")))
                .withMetricNamespace("microsoft.compute/virtualMachines").withMetricName("cpuusage")
                .withTimeGrain("PT1M").withAggregationType(MetricAggregationType.NONE).withDimension("nodename")
                .withDimensionFilter("node1"))
            .create();
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
