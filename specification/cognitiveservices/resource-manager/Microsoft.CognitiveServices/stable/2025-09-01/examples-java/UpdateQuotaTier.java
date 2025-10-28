
import com.azure.resourcemanager.cognitiveservices.models.QuotaTier;
import com.azure.resourcemanager.cognitiveservices.models.QuotaTierProperties;
import com.azure.resourcemanager.cognitiveservices.models.TierUpgradePolicy;

/**
 * Samples for QuotaTiers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * UpdateQuotaTier.json
     */
    /**
     * Sample code: Update the quota tier resource for a subscription.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateTheQuotaTierResourceForASubscription(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        QuotaTier resource
            = manager.quotaTiers().getWithResponse("default", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new QuotaTierProperties().withTierUpgradePolicy(TierUpgradePolicy.NO_AUTO_UPGRADE)).apply();
    }
}
