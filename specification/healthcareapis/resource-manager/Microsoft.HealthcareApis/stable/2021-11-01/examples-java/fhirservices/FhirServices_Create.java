import com.azure.resourcemanager.healthcareapis.models.FhirServiceAccessPolicyEntry;
import com.azure.resourcemanager.healthcareapis.models.FhirServiceAcrConfiguration;
import com.azure.resourcemanager.healthcareapis.models.FhirServiceAuthenticationConfiguration;
import com.azure.resourcemanager.healthcareapis.models.FhirServiceCorsConfiguration;
import com.azure.resourcemanager.healthcareapis.models.FhirServiceExportConfiguration;
import com.azure.resourcemanager.healthcareapis.models.FhirServiceKind;
import com.azure.resourcemanager.healthcareapis.models.ServiceManagedIdentityIdentity;
import com.azure.resourcemanager.healthcareapis.models.ServiceManagedIdentityType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for FhirServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Create.json
     */
    /**
     * Sample code: Create or update a Fhir Service.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void createOrUpdateAFhirService(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .fhirServices()
            .define("fhirservice1")
            .withExistingWorkspace("testRG", "workspace1")
            .withRegion("westus")
            .withTags(mapOf("additionalProp1", "string", "additionalProp2", "string", "additionalProp3", "string"))
            .withKind(FhirServiceKind.FHIR_R4)
            .withIdentity(new ServiceManagedIdentityIdentity().withType(ServiceManagedIdentityType.SYSTEM_ASSIGNED))
            .withAccessPolicies(
                Arrays
                    .asList(
                        new FhirServiceAccessPolicyEntry().withObjectId("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
                        new FhirServiceAccessPolicyEntry().withObjectId("5b307da8-43d4-492b-8b66-b0294ade872f")))
            .withAcrConfiguration(new FhirServiceAcrConfiguration().withLoginServers(Arrays.asList("test1.azurecr.io")))
            .withAuthenticationConfiguration(
                new FhirServiceAuthenticationConfiguration()
                    .withAuthority("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc")
                    .withAudience("https://azurehealthcareapis.com")
                    .withSmartProxyEnabled(true))
            .withCorsConfiguration(
                new FhirServiceCorsConfiguration()
                    .withOrigins(Arrays.asList("*"))
                    .withHeaders(Arrays.asList("*"))
                    .withMethods(Arrays.asList("DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT"))
                    .withMaxAge(1440)
                    .withAllowCredentials(false))
            .withExportConfiguration(
                new FhirServiceExportConfiguration().withStorageAccountName("existingStorageAccount"))
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
