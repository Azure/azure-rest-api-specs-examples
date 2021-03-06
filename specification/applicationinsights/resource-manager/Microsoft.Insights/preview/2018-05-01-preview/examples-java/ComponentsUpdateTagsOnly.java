import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.ApplicationInsightsComponent;
import java.util.HashMap;
import java.util.Map;

/** Samples for Components UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsUpdateTagsOnly.json
     */
    /**
     * Sample code: ComponentUpdateTagsOnly.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentUpdateTagsOnly(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        ApplicationInsightsComponent resource =
            manager
                .components()
                .getByResourceGroupWithResponse("my-resource-group", "my-component", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(
                mapOf(
                    "ApplicationGatewayType",
                    "Internal-Only",
                    "BillingEntity",
                    "Self",
                    "Color",
                    "AzureBlue",
                    "CustomField_01",
                    "Custom text in some random field named randomly",
                    "NodeType",
                    "Edge"))
            .apply();
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
