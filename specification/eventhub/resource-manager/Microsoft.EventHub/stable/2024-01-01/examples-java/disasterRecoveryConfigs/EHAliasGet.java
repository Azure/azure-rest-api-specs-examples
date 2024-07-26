
/**
 * Samples for DisasterRecoveryConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/disasterRecoveryConfigs/
     * EHAliasGet.json
     */
    /**
     * Sample code: EHAliasGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eHAliasGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getDisasterRecoveryConfigs().getWithResponse("exampleResourceGroup",
            "sdk-Namespace-8859", "sdk-DisasterRecovery-3814", com.azure.core.util.Context.NONE);
    }
}
