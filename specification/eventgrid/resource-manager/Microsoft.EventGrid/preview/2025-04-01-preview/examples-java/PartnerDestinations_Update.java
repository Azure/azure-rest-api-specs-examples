
import com.azure.resourcemanager.eventgrid.models.PartnerDestination;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PartnerDestinations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * PartnerDestinations_Update.json
     */
    /**
     * Sample code: PartnerDestinations_Update.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerDestinationsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        PartnerDestination resource = manager.partnerDestinations().getByResourceGroupWithResponse("examplerg",
            "examplePartnerDestinationName1", com.azure.core.util.Context.NONE).getValue();
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
