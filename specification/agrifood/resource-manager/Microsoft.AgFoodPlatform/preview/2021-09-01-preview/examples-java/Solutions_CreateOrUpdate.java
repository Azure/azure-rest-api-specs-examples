
import com.azure.resourcemanager.agrifood.models.SolutionProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Solutions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/
     * Solutions_CreateOrUpdate.json
     */
    /**
     * Sample code: Solutions_CreateOrUpdate.
     * 
     * @param manager Entry point to AgriFoodManager.
     */
    public static void solutionsCreateOrUpdate(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.solutions().define("abc.partner").withExistingFarmBeat("examples-rg", "examples-farmbeatsResourceName")
            .withProperties(new SolutionProperties().withSaasSubscriptionId("123").withSaasSubscriptionName("name")
                .withMarketplacePublisherId("publisherId").withPlanId("planId").withOfferId("offerId")
                .withTermId("termId").withAdditionalProperties(mapOf()))
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
