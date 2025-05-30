
/**
 * Samples for MarketplaceAgreements List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MarketplaceAgreements_List.
     * json
     */
    /**
     * Sample code: MarketplaceAgreements_List.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void marketplaceAgreementsList(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.marketplaceAgreements().list(com.azure.core.util.Context.NONE);
    }
}
