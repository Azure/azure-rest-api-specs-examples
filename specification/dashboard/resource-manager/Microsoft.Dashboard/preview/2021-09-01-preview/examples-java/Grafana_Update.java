import com.azure.core.util.Context;
import com.azure.resourcemanager.dashboard.models.ManagedGrafana;
import java.util.HashMap;
import java.util.Map;

/** Samples for Grafana Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_Update.json
     */
    /**
     * Sample code: Grafana_Update.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaUpdate(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        ManagedGrafana resource =
            manager
                .grafanas()
                .getByResourceGroupWithResponse("myResourceGroup", "myWorkspace", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("Environment", "Dev 2")).apply();
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
