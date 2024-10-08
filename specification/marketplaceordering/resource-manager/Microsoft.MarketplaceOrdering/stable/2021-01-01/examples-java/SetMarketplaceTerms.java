
import com.azure.resourcemanager.marketplaceordering.fluent.models.AgreementTermsInner;
import com.azure.resourcemanager.marketplaceordering.models.OfferType;
import java.time.OffsetDateTime;

/**
 * Samples for MarketplaceAgreements Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/
     * SetMarketplaceTerms.json
     */
    /**
     * Sample code: SetMarketplaceTerms.
     * 
     * @param manager Entry point to MarketplaceOrderingManager.
     */
    public static void
        setMarketplaceTerms(com.azure.resourcemanager.marketplaceordering.MarketplaceOrderingManager manager) {
        manager.marketplaceAgreements().createWithResponse(OfferType.VIRTUALMACHINE, "pubid", "offid", "planid",
            new AgreementTermsInner().withPublisher("pubid").withProduct("offid").withPlan("planid")
                .withLicenseTextLink("test.licenseLink").withPrivacyPolicyLink("test.privacyPolicyLink")
                .withMarketplaceTermsLink("test.marketplaceTermsLink")
                .withRetrieveDatetime(OffsetDateTime.parse("2017-08-15T11:33:07.12132Z"))
                .withSignature("ASDFSDAFWEFASDGWERLWER").withAccepted(false),
            com.azure.core.util.Context.NONE);
    }
}
