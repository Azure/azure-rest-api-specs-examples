import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.workloads.models.Monitor;
import com.azure.resourcemanager.workloads.models.UserAssignedServiceIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for Monitors Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_PatchTags_Delete.json
     */
    /**
     * Sample code: Delete Tags field of a SAP monitor.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void deleteTagsFieldOfASAPMonitor(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        Monitor resource =
            manager
                .monitors()
                .getByResourceGroupWithResponse("myResourceGroup", "mySapMonitor", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf())
            .withIdentity(new UserAssignedServiceIdentity().withType(ManagedServiceIdentityType.NONE))
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
