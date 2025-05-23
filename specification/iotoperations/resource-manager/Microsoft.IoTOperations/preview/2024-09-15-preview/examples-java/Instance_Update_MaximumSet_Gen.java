
import com.azure.resourcemanager.iotoperations.models.InstanceResource;
import com.azure.resourcemanager.iotoperations.models.ManagedServiceIdentity;
import com.azure.resourcemanager.iotoperations.models.ManagedServiceIdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Instance Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Instance_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instance_Update.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void instanceUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        InstanceResource resource = manager.instances()
            .getByResourceGroupWithResponse("rgiotoperations", "aio-instance", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf()).withIdentity(
            new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE).withUserAssignedIdentities(mapOf()))
            .apply();
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
