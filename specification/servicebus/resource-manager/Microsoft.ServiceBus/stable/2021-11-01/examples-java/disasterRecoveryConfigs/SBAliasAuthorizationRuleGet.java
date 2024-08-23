
/**
 * Samples for DisasterRecoveryConfigs GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasAuthorizationRuleGet.json
     */
    /**
     * Sample code: DisasterRecoveryConfigsAuthorizationRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        disasterRecoveryConfigsAuthorizationRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs()
            .getAuthorizationRuleWithResponse("exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4879",
                "sdk-Authrules-4879", com.azure.core.util.Context.NONE);
    }
}
