
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs GetAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/
     * EHAliasAuthorizationRuleGet.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceAuthorizationRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getDisasterRecoveryConfigs().getAuthorizationRuleWithResponse(
            "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879",
            Context.NONE);
    }
}
