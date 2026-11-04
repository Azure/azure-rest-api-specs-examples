
import com.azure.resourcemanager.koreaiartemis.models.CompanyDetails;
import com.azure.resourcemanager.koreaiartemis.models.DefaultWorkspaceProperties;
import com.azure.resourcemanager.koreaiartemis.models.MarketplaceDetails;
import com.azure.resourcemanager.koreaiartemis.models.OfferDetails;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationPartnerProperties;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationProperties;
import com.azure.resourcemanager.koreaiartemis.models.Product;
import com.azure.resourcemanager.koreaiartemis.models.UserDetails;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Create.json
     */
    /**
     * Sample code: Organizations_Create.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void organizationsCreate(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().define("contoso-org").withRegion("eastus2euap").withExistingResourceGroup("contoso-rg")
            .withTags(mapOf("environment", "production"))
            .withProperties(new OrganizationProperties().withPartnerProperties(new OrganizationPartnerProperties()
                .withOrganizationName("Contoso AI Platform").withBillingEmail("billing@contoso.com")
                .withOwnerEmail("admin@contoso.com")
                .withDefaultWorkspace(new DefaultWorkspaceProperties().withName("my first workspace")
                    .withProduct(Product.AIFORWORK)
                    .withMarketplaceDetails(Arrays.asList(new MarketplaceDetails().withSubscriptionId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas")
                        .withOfferDetails(new OfferDetails().withPublisherId("kore-ai").withOfferId("kore-ai-platform")
                            .withPlanId("business-monthly").withPlanName("Business Monthly").withTermUnit("P1M")
                            .withTermId("1m"))))))
                .withUser(
                    new UserDetails().withFirstName("Jane").withLastName("Doe").withEmailAddress("jane.doe@contoso.com")
                        .withUpn("jane.doe@contoso.onmicrosoft.com").withPhoneNumber("+1-425-555-0100"))
                .withCompanyDetails(new CompanyDetails().withCompanyName("Contoso Ltd.").withCountry("US")
                    .withDomain("contoso.com").withBusinessType("Technology").withNumberOfEmployees(500)))
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
