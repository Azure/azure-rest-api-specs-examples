import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.workloads.models.PatchResourceRequestBodyIdentity;
import com.azure.resourcemanager.workloads.models.PhpWorkloadResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for PhpWorkloads Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/PhpWorkloads_Update.json
     */
    /**
     * Sample code: Workloads.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void workloads(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        PhpWorkloadResource resource =
            manager.phpWorkloads().getByResourceGroupWithResponse("test-rg", "wp39", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("tag_name", "tag_value"))
            .withIdentity(new PatchResourceRequestBodyIdentity().withType(ManagedServiceIdentityType.NONE))
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
