
/**
 * Samples for DisasterRecoveryConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasDelete.json
     */
    /**
     * Sample code: EHAliasDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHAliasDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().deleteWithResponse("exampleResourceGroup",
            "sdk-Namespace-5849", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
