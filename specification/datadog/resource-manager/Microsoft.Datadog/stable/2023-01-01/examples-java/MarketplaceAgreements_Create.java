/** Samples for MarketplaceAgreements CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MarketplaceAgreements_Create.json
     */
    /**
     * Sample code: MarketplaceAgreements_CreateOrUpdate.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void marketplaceAgreementsCreateOrUpdate(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.marketplaceAgreements().createOrUpdateWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
