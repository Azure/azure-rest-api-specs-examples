
/**
 * Samples for DisasterRecoveryConfigs ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasAuthorizationRuleListKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleListKey.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceAuthorizationRuleListKey(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().listKeysWithResponse("exampleResourceGroup",
            "sdk-Namespace-2702", "sdk-DisasterRecovery-4047", "sdk-Authrules-1746", com.azure.core.util.Context.NONE);
    }
}
