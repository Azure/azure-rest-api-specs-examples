
/**
 * Samples for DisasterRecoveryConfigs FailOver.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasFailOver.json
     */
    /**
     * Sample code: EHAliasFailOver.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHAliasFailOver(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().failOverWithResponse("exampleResourceGroup",
            "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
