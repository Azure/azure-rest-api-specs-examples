
/**
 * Samples for DisasterRecoveryConfigs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasList.json
     */
    /**
     * Sample code: EHAliasList.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHAliasList(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().list("exampleResourceGroup", "sdk-Namespace-8859",
            com.azure.core.util.Context.NONE);
    }
}
