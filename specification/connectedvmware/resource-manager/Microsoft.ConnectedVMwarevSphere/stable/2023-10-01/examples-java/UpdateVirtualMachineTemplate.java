
import com.azure.resourcemanager.connectedvmware.models.VirtualMachineTemplate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineTemplates Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/
     * UpdateVirtualMachineTemplate.json
     */
    /**
     * Sample code: UpdateVirtualMachineTemplate.
     * 
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void
        updateVirtualMachineTemplate(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        VirtualMachineTemplate resource = manager.virtualMachineTemplates()
            .getByResourceGroupWithResponse("testrg", "WebFrontEndTemplate", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
