
import com.azure.resourcemanager.koreaiartemis.models.OrganizationPartnerPropertiesUpdate;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationPropertiesUpdate;
import com.azure.resourcemanager.koreaiartemis.models.OrganizationResource;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Update.json
     */
    /**
     * Sample code: Organizations_Update.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void organizationsUpdate(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        OrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("contoso-rg", "contoso-org", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new OrganizationPropertiesUpdate().withPartnerProperties(
                new OrganizationPartnerPropertiesUpdate().withOrganizationName("Contoso AI Platform (Updated)")
                    .withBillingEmail("new-billing@contoso.com").withOwnerEmail("admin@contoso.com")))
            .apply();
    }
}
