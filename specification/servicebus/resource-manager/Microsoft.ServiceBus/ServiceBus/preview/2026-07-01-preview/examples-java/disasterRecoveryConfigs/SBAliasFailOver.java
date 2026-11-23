
/**
 * Samples for DisasterRecoveryConfigs FailOver.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasFailOver.json
     */
    /**
     * Sample code: SBAliasFailOver.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void sBAliasFailOver(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().failOverWithResponse("ardsouzatestRG",
            "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", null, com.azure.core.util.Context.NONE);
    }
}
