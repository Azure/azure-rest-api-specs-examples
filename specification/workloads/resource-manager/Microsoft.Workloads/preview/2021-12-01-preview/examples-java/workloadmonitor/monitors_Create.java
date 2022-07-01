import com.azure.resourcemanager.workloads.models.ManagedRGConfiguration;
import com.azure.resourcemanager.workloads.models.RoutingPreference;
import java.util.HashMap;
import java.util.Map;

/** Samples for Monitors Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_Create.json
     */
    /**
     * Sample code: Create a SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void createASAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .monitors()
            .define("mySapMonitor")
            .withRegion("westus")
            .withExistingResourceGroup("myResourceGroup")
            .withTags(mapOf("key", "value"))
            .withAppLocation("westus")
            .withRoutingPreference(RoutingPreference.ROUTE_ALL)
            .withManagedResourceGroupConfiguration(new ManagedRGConfiguration().withName("myManagedRg"))
            .withLogAnalyticsWorkspaceArmId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace")
            .withMonitorSubnet(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet")
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
