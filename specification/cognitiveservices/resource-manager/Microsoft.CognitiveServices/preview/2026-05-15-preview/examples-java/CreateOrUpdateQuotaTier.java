
import com.azure.resourcemanager.cognitiveservices.models.QuotaTierProperties;
import com.azure.resourcemanager.cognitiveservices.models.TierUpgradePolicy;

/**
 * Samples for QuotaTiers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/CreateOrUpdateQuotaTier.json
     */
    /**
     * Sample code: Update the quota tier resource for a subscription.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateTheQuotaTierResourceForASubscription(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.quotaTiers().define("default")
            .withProperties(new QuotaTierProperties().withTierUpgradePolicy(TierUpgradePolicy.NO_AUTO_UPGRADE))
            .create();
    }
}
