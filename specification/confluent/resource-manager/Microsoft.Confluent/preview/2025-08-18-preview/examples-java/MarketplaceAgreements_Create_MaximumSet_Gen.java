
import com.azure.resourcemanager.confluent.fluent.models.ConfluentAgreementResourceInner;
import java.time.OffsetDateTime;

/**
 * Samples for MarketplaceAgreements Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/MarketplaceAgreements_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: Create Confluent Marketplace agreement in the subscription. (Maximumset).
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void createConfluentMarketplaceAgreementInTheSubscriptionMaximumset(
        com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements()
            .createWithResponse(
                new ConfluentAgreementResourceInner().withPublisher("cxcrrfggvdmvcchohkyatlvbpyy")
                    .withProduct("ogusipjbwihlwbfivdbjfuvoqwija").withPlan("vgphlikczel")
                    .withLicenseTextLink("ztckliskduxmcluia").withPrivacyPolicyLink("wwvlrlfhzmvfjgimkhkqcaxn")
                    .withRetrieveDatetime(OffsetDateTime.parse("2025-08-18T11:10:31.028Z"))
                    .withSignature("cfdxpybzzsrgcdtebmqzzskxfiool").withAccepted(true),
                com.azure.core.util.Context.NONE);
    }
}
