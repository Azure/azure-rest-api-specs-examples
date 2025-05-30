
import com.azure.resourcemanager.devspaces.models.Controller;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Controllers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersUpdate_example
     * .json
     */
    /**
     * Sample code: ControllersUpdate.
     * 
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersUpdate(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        Controller resource = manager.controllers()
            .getByResourceGroupWithResponse("myResourceGroup", "myControllerResource", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key", "fakeTokenPlaceholder"))
            .withTargetContainerHostCredentialsBase64("QmFzZTY0IEVuY29kZWQgVmFsdWUK").apply();
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
