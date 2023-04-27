import com.azure.resourcemanager.resourceconnector.models.Appliance;
import java.util.HashMap;
import java.util.Map;

/** Samples for Appliances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesPatch.json
     */
    /**
     * Sample code: Update Appliance.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void updateAppliance(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        Appliance resource =
            manager
                .appliances()
                .getByResourceGroupWithResponse("testresourcegroup", "appliance01", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key", "value")).apply();
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
