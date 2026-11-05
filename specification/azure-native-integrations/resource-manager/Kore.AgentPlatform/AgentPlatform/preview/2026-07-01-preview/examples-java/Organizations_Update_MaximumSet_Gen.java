
import com.azure.resourcemanager.koreaiartemis.models.CompanyDetails;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationPartnerPropertiesUpdate;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationPropertiesUpdate;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationResource;
import com.azure.resourcemanager.koreaiartemis.models.UserDetails;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        organizationsUpdateMaximumSetGen(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        OrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("contoso-rg", "contoso-org", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new OrganizationPropertiesUpdate()
                .withPartnerProperties(
                    new OrganizationPartnerPropertiesUpdate().withOrganizationName("Contoso AI Platform (Updated)")
                        .withBillingEmail("new-billing@contoso.com").withOwnerEmail("admin@contoso.com"))
                .withUser(new UserDetails().withFirstName("Jane").withLastName("Doe")
                    .withEmailAddress("jane.alternate@contoso.com").withUpn("jane.doe@contoso.onmicrosoft.com")
                    .withPhoneNumber("+1-425-555-0100").withAlternateEmail("jane.alternate@contoso.com")
                    .withAddress("1 Microsoft Way, Redmond, WA 98052"))
                .withCompanyDetails(new CompanyDetails().withCompanyName("Contoso Ltd.")
                    .withOfficeAddress("1 Microsoft Way, Redmond, WA 98052").withCountry("US").withDomain("contoso.com")
                    .withBusinessType("Technology").withNumberOfEmployees(28)))
            .apply();
    }
}
