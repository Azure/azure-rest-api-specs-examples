
/**
 * Samples for DisasterRecoveryConfigs GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/EHAliasAuthorizationRuleGet.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceAuthorizationRuleGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().getAuthorizationRuleWithResponse("exampleResourceGroup",
            "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879", com.azure.core.util.Context.NONE);
    }
}
