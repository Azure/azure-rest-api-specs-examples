
/**
 * Samples for DisasterRecoveryConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasGet.json
     */
    /**
     * Sample code: EHAliasGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHAliasGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().getWithResponse("exampleResourceGroup",
            "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
