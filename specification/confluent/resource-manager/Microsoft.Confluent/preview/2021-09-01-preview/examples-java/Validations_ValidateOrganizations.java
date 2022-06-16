import com.azure.core.util.Context;
import com.azure.resourcemanager.confluent.fluent.models.OrganizationResourceInner;
import com.azure.resourcemanager.confluent.models.OfferDetail;
import com.azure.resourcemanager.confluent.models.UserDetail;
import java.util.HashMap;
import java.util.Map;

/** Samples for Validations ValidateOrganization. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Validations_ValidateOrganizations.json
     */
    /**
     * Sample code: Validations_ValidateOrganizations.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void validationsValidateOrganizations(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager
            .validations()
            .validateOrganizationWithResponse(
                "myResourceGroup",
                "myOrganization",
                new OrganizationResourceInner()
                    .withLocation("West US")
                    .withTags(mapOf("Environment", "Dev"))
                    .withOfferDetail(
                        new OfferDetail()
                            .withPublisherId("string")
                            .withId("string")
                            .withPlanId("string")
                            .withPlanName("string")
                            .withTermUnit("string"))
                    .withUserDetail(
                        new UserDetail()
                            .withFirstName("string")
                            .withLastName("string")
                            .withEmailAddress("abc@microsoft.com")),
                Context.NONE);
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
