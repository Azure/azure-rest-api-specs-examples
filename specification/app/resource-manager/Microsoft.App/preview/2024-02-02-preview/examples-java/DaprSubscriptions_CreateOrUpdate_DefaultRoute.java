
import com.azure.resourcemanager.appcontainers.models.DaprSubscriptionRoutes;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DaprSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * DaprSubscriptions_CreateOrUpdate_DefaultRoute.json
     */
    /**
     * Sample code: Create or update dapr subscription with default route only.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateDaprSubscriptionWithDefaultRouteOnly(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprSubscriptions().define("mysubscription")
            .withExistingManagedEnvironment("examplerg", "myenvironment").withPubsubName("mypubsubcomponent")
            .withTopic("inventory").withRoutes(new DaprSubscriptionRoutes().withDefaultProperty("/products")).create();
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
