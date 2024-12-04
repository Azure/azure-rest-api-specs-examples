
import com.azure.resourcemanager.eventgrid.models.Partner;
import java.time.OffsetDateTime;
import java.util.UUID;

/**
 * Samples for PartnerConfigurations AuthorizePartner.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * PartnerConfigurations_AuthorizePartner.json
     */
    /**
     * Sample code: PartnerConfigurations_AuthorizePartner.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        partnerConfigurationsAuthorizePartner(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerConfigurations().authorizePartnerWithResponse("examplerg",
            new Partner().withPartnerRegistrationImmutableId(UUID.fromString("941892bc-f5d0-4d1c-8fb5-477570fc2b71"))
                .withPartnerName("Contoso.Finance").withAuthorizationExpirationTimeInUtc(
                    OffsetDateTime.parse("2022-01-28T01:20:55.142Z")),
            com.azure.core.util.Context.NONE);
    }
}
