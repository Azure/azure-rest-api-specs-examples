
import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.fluent.models.RuleInner;
import com.azure.resourcemanager.servicebus.models.CorrelationFilter;
import com.azure.resourcemanager.servicebus.models.FilterType;
import java.util.HashMap;
import java.util.Map;

/** Samples for Rules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/
     * RuleCreate_CorrelationFilter.json
     */
    /**
     * Sample code: RulesCreateCorrelationFilter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesCreateCorrelationFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getRules().createOrUpdateWithResponse(
            "resourceGroupName", "sdk-Namespace-1319", "sdk-Topics-2081", "sdk-Subscriptions-8691", "sdk-Rules-6571",
            new RuleInner().withFilterType(FilterType.CORRELATION_FILTER).withCorrelationFilter(
                new CorrelationFilter().withProperties(mapOf("topicHint", "Crop"))),
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
