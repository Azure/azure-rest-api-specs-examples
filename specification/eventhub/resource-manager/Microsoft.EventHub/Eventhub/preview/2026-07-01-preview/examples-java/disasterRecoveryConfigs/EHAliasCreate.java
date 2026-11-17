
import com.azure.resourcemanager.eventhubs.fluent.models.ArmDisasterRecoveryInner;

/**
 * Samples for DisasterRecoveryConfigs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasCreate.json
     */
    /**
     * Sample code: EHAliasCreate.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHAliasCreate(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().createOrUpdateWithResponse("exampleResourceGroup",
            "sdk-Namespace-8859", "sdk-DisasterRecovery-3814",
            new ArmDisasterRecoveryInner().withPartnerNamespace("sdk-Namespace-37"), com.azure.core.util.Context.NONE);
    }
}
