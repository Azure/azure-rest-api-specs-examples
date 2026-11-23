
import com.azure.resourcemanager.servicebus.fluent.models.RuleInner;
import com.azure.resourcemanager.servicebus.models.FilterType;
import com.azure.resourcemanager.servicebus.models.SqlFilter;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Rules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Rules/RuleCreate_SqlFilter.json
     */
    /**
     * Sample code: RulesCreateSqlFilter.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void rulesCreateSqlFilter(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getRules()
            .createOrUpdateWithResponse("resourceGroupName", "sdk-Namespace-1319", "sdk-Topics-2081",
                "sdk-Subscriptions-8691", "sdk-Rules-6571", new RuleInner().withFilterType(FilterType.SQL_FILTER)
                    .withSqlFilter(new SqlFilter().withSqlExpression("myproperty=test")),
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
