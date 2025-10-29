
import com.azure.resourcemanager.cognitiveservices.models.QuotaTierProperties;
import com.azure.resourcemanager.cognitiveservices.models.TierUpgradePolicy;

/**
 * Samples for QuotaTiers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * CreateOrUpdateQuotaTier.json
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
