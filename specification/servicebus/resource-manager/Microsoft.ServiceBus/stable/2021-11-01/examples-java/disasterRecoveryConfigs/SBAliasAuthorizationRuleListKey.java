
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs ListKeys. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasAuthorizationRuleListKey.json
     */
    /**
     * Sample code: DisasterRecoveryConfigsAuthorizationRuleListKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        disasterRecoveryConfigsAuthorizationRuleListKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs().listKeysWithResponse(
            "exampleResourceGroup", "sdk-Namespace-2702", "sdk-DisasterRecovery-4047", "sdk-Authrules-1746",
            Context.NONE);
    }
}
