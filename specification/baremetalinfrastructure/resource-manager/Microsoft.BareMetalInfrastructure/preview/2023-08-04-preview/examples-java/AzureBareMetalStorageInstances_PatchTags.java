
import com.azure.resourcemanager.baremetalinfrastructure.models.AzureBareMetalStorageInstance;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AzureBareMetalStorageInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-
     * preview/examples/AzureBareMetalStorageInstances_PatchTags.json
     */
    /**
     * Sample code: Update Tags field of an AzureBareMetalStorage instance.
     * 
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void updateTagsFieldOfAnAzureBareMetalStorageInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        AzureBareMetalStorageInstance resource = manager.azureBareMetalStorageInstances()
            .getByResourceGroupWithResponse("myResourceGroup", "myABMSInstance", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("testkey", "fakeTokenPlaceholder")).apply();
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
