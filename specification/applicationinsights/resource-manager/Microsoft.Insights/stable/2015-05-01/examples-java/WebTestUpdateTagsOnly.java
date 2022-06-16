import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.WebTest;
import java.util.HashMap;
import java.util.Map;

/** Samples for WebTests UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestUpdateTagsOnly.json
     */
    /**
     * Sample code: webTestUpdateTags.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestUpdateTags(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        WebTest resource =
            manager
                .webTests()
                .getByResourceGroupWithResponse("my-resource-group", "my-webtest-my-component", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(
                mapOf(
                    "Color",
                    "AzureBlue",
                    "CustomField-01",
                    "This is a random value",
                    "SystemType",
                    "A08",
                    "hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component",
                    "Resource",
                    "hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Web/sites/mytestwebapp",
                    "Resource"))
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
