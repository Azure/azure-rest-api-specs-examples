
/**
 * Samples for DisasterRecoveryConfigs ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/disasterRecoveryConfigs/SBAliasAuthorizationRuleListAll.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleListAll.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpaceAuthorizationRuleListAll(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getDisasterRecoveryConfigs().listAuthorizationRules("exampleResourceGroup",
            "sdk-Namespace-9080", "sdk-DisasterRecovery-4047", com.azure.core.util.Context.NONE);
    }
}
