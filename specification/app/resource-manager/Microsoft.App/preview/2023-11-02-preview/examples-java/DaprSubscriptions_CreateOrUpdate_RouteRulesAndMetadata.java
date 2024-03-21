
import com.azure.resourcemanager.appcontainers.models.DaprSubscriptionRouteRule;
import com.azure.resourcemanager.appcontainers.models.DaprSubscriptionRoutes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DaprSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/
     * DaprSubscriptions_CreateOrUpdate_RouteRulesAndMetadata.json
     */
    /**
     * Sample code: Create or update dapr subscription with route rules and metadata.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateDaprSubscriptionWithRouteRulesAndMetadata(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprSubscriptions().define("mysubscription")
            .withExistingManagedEnvironment("examplerg", "myenvironment").withPubsubName("mypubsubcomponent")
            .withTopic("inventory")
            .withRoutes(new DaprSubscriptionRoutes()
                .withRules(Arrays.asList(
                    new DaprSubscriptionRouteRule().withMatch("event.type == 'widget'").withPath("/widgets"),
                    new DaprSubscriptionRouteRule().withMatch("event.type == 'gadget'").withPath("/gadgets")))
                .withDefaultProperty("/products"))
            .withMetadata(mapOf("foo", "bar", "hello", "world")).create();
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
