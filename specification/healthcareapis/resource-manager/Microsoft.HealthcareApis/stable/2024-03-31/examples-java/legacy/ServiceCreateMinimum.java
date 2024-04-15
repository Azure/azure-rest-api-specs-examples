
import com.azure.resourcemanager.healthcareapis.models.Kind;
import com.azure.resourcemanager.healthcareapis.models.ServiceAccessPolicyEntry;
import com.azure.resourcemanager.healthcareapis.models.ServicesProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceCreateMinimum.json
     */
    /**
     * Sample code: Create or Update a service with minimum parameters.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void createOrUpdateAServiceWithMinimumParameters(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().define("service2").withRegion("westus2").withExistingResourceGroup("rg1")
            .withKind(Kind.FHIR_R4).withTags(mapOf())
            .withProperties(new ServicesProperties().withAccessPolicies(
                Arrays.asList(new ServiceAccessPolicyEntry().withObjectId("c487e7d1-3210-41a3-8ccc-e9372b78da47"))))
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
