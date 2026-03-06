
/**
 * Samples for RateCard Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-06-01-preview/GetRateCard.json
     */
    /**
     * Sample code: GetRateCard.
     * 
     * @param manager Entry point to CommerceManager.
     */
    public static void getRateCard(com.azure.resourcemanager.commerce.CommerceManager manager) {
        manager.rateCards().getWithResponse(
            "OfferDurableId eq 'MS-AZR-0003P' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'",
            com.azure.core.util.Context.NONE);
    }
}
