
import com.azure.resourcemanager.customerinsights.models.PropertyDefinition;
import java.util.Arrays;

/**
 * Samples for Interactions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/
     * InteractionsCreateOrUpdate.json
     */
    /**
     * Sample code: Interactions_CreateOrUpdate.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void
        interactionsCreateOrUpdate(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.interactions().define("TestProfileType396").withExistingHub("TestHubRG", "sdkTestHub")
            .withIdPropertyNames(Arrays.asList("TestInteractionType6358"))
            .withPrimaryParticipantProfilePropertyName("profile1").withApiEntitySetName("TestInteractionType6358")
            .withFields(Arrays.asList(
                new PropertyDefinition().withFieldName("TestInteractionType6358").withFieldType("Edm.String")
                    .withIsArray(false).withIsRequired(true),
                new PropertyDefinition().withFieldName("profile1").withFieldType("Edm.String")))
            .withSmallImage("\\\\Images\\\\smallImage").withMediumImage("\\\\Images\\\\MediumImage")
            .withLargeImage("\\\\Images\\\\LargeImage").create();
    }
}
