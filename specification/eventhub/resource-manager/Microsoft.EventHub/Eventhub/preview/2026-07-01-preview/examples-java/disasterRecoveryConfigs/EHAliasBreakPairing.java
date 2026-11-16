
/**
 * Samples for DisasterRecoveryConfigs BreakPairing.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasBreakPairing.json
     */
    /**
     * Sample code: EHAliasBreakPairing.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHAliasBreakPairing(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().breakPairingWithResponse("exampleResourceGroup",
            "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
