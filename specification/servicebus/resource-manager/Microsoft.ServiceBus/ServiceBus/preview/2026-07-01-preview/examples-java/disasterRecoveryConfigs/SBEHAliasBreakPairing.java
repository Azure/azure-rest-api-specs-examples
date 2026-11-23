
/**
 * Samples for DisasterRecoveryConfigs BreakPairing.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBEHAliasBreakPairing.json
     */
    /**
     * Sample code: SBEHAliasBreakPairing.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void sBEHAliasBreakPairing(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().breakPairingWithResponse("ardsouzatestRG",
            "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
