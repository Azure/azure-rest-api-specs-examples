import java.util.HashMap;
import java.util.Map;

/** Samples for DevCenters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevCenters_Create.json
     */
    /**
     * Sample code: DevCenters_Create.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersCreate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .devCenters()
            .define("Contoso")
            .withRegion("centralus")
            .withExistingResourceGroup("rg1")
            .withTags(mapOf("CostCode", "fakeTokenPlaceholder"))
            .withDisplayName("ContosoDevCenter")
            .create();
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
