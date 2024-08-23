
/**
 * Samples for DisasterRecoveryConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasDelete.json
     */
    /**
     * Sample code: SBAliasDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sBAliasDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs().deleteWithResponse(
            "SouthCentralUS", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
