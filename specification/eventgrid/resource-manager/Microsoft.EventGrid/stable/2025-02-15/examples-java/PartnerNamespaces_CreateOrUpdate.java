
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PartnerNamespaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerNamespaces_CreateOrUpdate.json
     */
    /**
     * Sample code: PartnerNamespaces_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerNamespaces().define("examplePartnerNamespaceName1").withRegion("westus")
            .withExistingResourceGroup("examplerg").withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withPartnerRegistrationFullyQualifiedId(
                "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerRegistrations/ContosoCorpAccount1")
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
