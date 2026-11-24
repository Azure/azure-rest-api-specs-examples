
/**
 * Samples for DisasterRecoveryConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasDelete.json
     */
    /**
     * Sample code: SBAliasDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void sBAliasDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().deleteWithResponse("SouthCentralUS", "sdk-Namespace-8860",
            "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
