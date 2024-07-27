
/**
 * Samples for DisasterRecoveryConfigs ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/disasterRecoveryConfigs/
     * EHAliasAuthorizationRuleListAll.json
     */
    /**
     * Sample code: ListAuthorizationRules.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuthorizationRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getDisasterRecoveryConfigs().listAuthorizationRules(
            "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4047",
            com.azure.core.util.Context.NONE);
    }
}
