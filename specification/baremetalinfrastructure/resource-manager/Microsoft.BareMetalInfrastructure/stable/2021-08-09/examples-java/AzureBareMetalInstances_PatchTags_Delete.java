import com.azure.core.util.Context;
import com.azure.resourcemanager.baremetalinfrastructure.models.Tags;
import java.util.HashMap;
import java.util.Map;

/** Samples for AzureBareMetalInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_PatchTags_Delete.json
     */
    /**
     * Sample code: Delete Tags field of an AzureBareMetal instance.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void deleteTagsFieldOfAnAzureBareMetalInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager
            .azureBareMetalInstances()
            .updateWithResponse("myResourceGroup", "myABMInstance", new Tags().withTags(mapOf()), Context.NONE);
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
