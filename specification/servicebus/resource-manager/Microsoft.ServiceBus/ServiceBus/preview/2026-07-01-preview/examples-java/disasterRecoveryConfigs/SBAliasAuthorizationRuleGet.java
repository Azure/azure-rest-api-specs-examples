
/**
 * Samples for DisasterRecoveryConfigs GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
     */
    /**
     * Sample code: DisasterRecoveryConfigsAuthorizationRuleGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        disasterRecoveryConfigsAuthorizationRuleGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().getAuthorizationRuleWithResponse("exampleResourceGroup",
            "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879", com.azure.core.util.Context.NONE);
    }
}
