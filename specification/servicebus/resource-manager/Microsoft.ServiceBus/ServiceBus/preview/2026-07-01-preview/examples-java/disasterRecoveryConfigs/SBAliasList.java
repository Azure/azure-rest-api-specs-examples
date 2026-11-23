
/**
 * Samples for DisasterRecoveryConfigs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasList.json
     */
    /**
     * Sample code: SBAliasList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void sBAliasList(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().list("ardsouzatestRG", "sdk-Namespace-8860",
            com.azure.core.util.Context.NONE);
    }
}
