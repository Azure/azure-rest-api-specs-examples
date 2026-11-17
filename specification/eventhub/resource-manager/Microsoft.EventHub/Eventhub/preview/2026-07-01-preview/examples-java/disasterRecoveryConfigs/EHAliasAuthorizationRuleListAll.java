
/**
 * Samples for DisasterRecoveryConfigs ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasAuthorizationRuleListAll.json
     */
    /**
     * Sample code: ListAuthorizationRules.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void listAuthorizationRules(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().listAuthorizationRules("exampleResourceGroup",
            "sdk-Namespace-9080", "sdk-DisasterRecovery-4047", com.azure.core.util.Context.NONE);
    }
}
