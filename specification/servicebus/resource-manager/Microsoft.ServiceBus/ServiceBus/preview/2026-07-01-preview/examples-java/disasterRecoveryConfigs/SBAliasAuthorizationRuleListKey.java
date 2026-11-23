
/**
 * Samples for DisasterRecoveryConfigs ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasAuthorizationRuleListKey.json
     */
    /**
     * Sample code: DisasterRecoveryConfigsAuthorizationRuleListKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void disasterRecoveryConfigsAuthorizationRuleListKey(
        com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().listKeysWithResponse("exampleResourceGroup",
            "sdk-Namespace-2702", "sdk-DisasterRecovery-4047", "sdk-Authrules-1746", com.azure.core.util.Context.NONE);
    }
}
