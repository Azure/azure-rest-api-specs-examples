import java.util.HashMap;
import java.util.Map;

/** Samples for PartnerRegistrations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerRegistrations_CreateOrUpdate.json
     */
    /**
     * Sample code: PartnerRegistrations_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerRegistrations()
            .define("examplePartnerRegistrationName1")
            .withRegion("global")
            .withExistingResourceGroup("examplerg")
            .withTags(mapOf("key1", "value1", "key2", "Value2", "key3", "Value3"))
            .create();
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
