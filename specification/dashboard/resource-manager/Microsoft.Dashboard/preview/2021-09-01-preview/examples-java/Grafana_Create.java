import com.azure.resourcemanager.dashboard.models.IdentityType;
import com.azure.resourcemanager.dashboard.models.ManagedGrafanaProperties;
import com.azure.resourcemanager.dashboard.models.ManagedIdentity;
import com.azure.resourcemanager.dashboard.models.ProvisioningState;
import com.azure.resourcemanager.dashboard.models.ResourceSku;
import com.azure.resourcemanager.dashboard.models.ZoneRedundancy;
import java.util.HashMap;
import java.util.Map;

/** Samples for Grafana Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_Create.json
     */
    /**
     * Sample code: Grafana_Create.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaCreate(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager
            .grafanas()
            .define("myWorkspace")
            .withRegion("West US")
            .withExistingResourceGroup("myResourceGroup")
            .withTags(mapOf("Environment", "Dev"))
            .withSku(new ResourceSku().withName("Standard"))
            .withProperties(
                new ManagedGrafanaProperties()
                    .withProvisioningState(ProvisioningState.ACCEPTED)
                    .withZoneRedundancy(ZoneRedundancy.ENABLED))
            .withIdentity(new ManagedIdentity().withType(IdentityType.SYSTEM_ASSIGNED))
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
