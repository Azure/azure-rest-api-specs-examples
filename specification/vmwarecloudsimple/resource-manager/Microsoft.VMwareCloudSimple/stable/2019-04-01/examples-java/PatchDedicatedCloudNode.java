
import com.azure.resourcemanager.vmwarecloudsimple.models.DedicatedCloudNode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DedicatedCloudNodes Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/
     * PatchDedicatedCloudNode.json
     */
    /**
     * Sample code: PatchDedicatedCloudNode.
     * 
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void
        patchDedicatedCloudNode(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        DedicatedCloudNode resource = manager.dedicatedCloudNodes()
            .getByResourceGroupWithResponse("myResourceGroup", "myNode", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("myTag", "tagValue")).apply();
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
