
/**
 * Samples for DisasterRecoveryConfigs FailOver.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasFailOver.json
     */
    /**
     * Sample code: SBAliasFailOver.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sBAliasFailOver(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs().failOverWithResponse(
            "ardsouzatestRG", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", null,
            com.azure.core.util.Context.NONE);
    }
}
