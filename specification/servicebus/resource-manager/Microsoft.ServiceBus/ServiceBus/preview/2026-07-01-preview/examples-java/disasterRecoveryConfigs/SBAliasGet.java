
/**
 * Samples for DisasterRecoveryConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasGet.json
     */
    /**
     * Sample code: SBAliasGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void sBAliasGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().getWithResponse("ardsouzatestRG", "sdk-Namespace-8860",
            "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
