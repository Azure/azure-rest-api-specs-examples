
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedDashboards Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Dashboard_Create.json
     */
    /**
     * Sample code: Dashboard_Create.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void dashboardCreate(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedDashboards().define("myDashboard").withRegion("West US")
            .withExistingResourceGroup("myResourceGroup").withTags(mapOf("Environment", "Dev")).create();
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
