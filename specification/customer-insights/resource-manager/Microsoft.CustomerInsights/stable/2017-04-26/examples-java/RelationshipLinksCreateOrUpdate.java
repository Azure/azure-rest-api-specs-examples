
import com.azure.resourcemanager.customerinsights.models.ParticipantProfilePropertyReference;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RelationshipLinks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/
     * RelationshipLinksCreateOrUpdate.json
     */
    /**
     * Sample code: RelationshipLinks_CreateOrUpdate.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void
        relationshipLinksCreateOrUpdate(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.relationshipLinks().define("Somelink").withExistingHub("TestHubRG", "sdkTestHub")
            .withDisplayName(mapOf("en-us", "Link DisplayName")).withDescription(mapOf("en-us", "Link Description"))
            .withInteractionType("testInteraction4332")
            .withProfilePropertyReferences(Arrays.asList(new ParticipantProfilePropertyReference()
                .withInteractionPropertyName("profile1").withProfilePropertyName("ProfileId")))
            .withRelatedProfilePropertyReferences(Arrays.asList(new ParticipantProfilePropertyReference()
                .withInteractionPropertyName("profile1").withProfilePropertyName("ProfileId")))
            .withRelationshipName("testProfile2326994").create();
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
