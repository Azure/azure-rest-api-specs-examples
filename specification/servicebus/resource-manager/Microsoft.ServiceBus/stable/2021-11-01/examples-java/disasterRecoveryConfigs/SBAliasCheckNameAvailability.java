
import com.azure.resourcemanager.servicebus.models.CheckNameAvailability;

/**
 * Samples for DisasterRecoveryConfigs CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs
     * /SBAliasCheckNameAvailability.json
     */
    /**
     * Sample code: AliasNameAvailability.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aliasNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getDisasterRecoveryConfigs()
            .checkNameAvailabilityWithResponse("exampleResourceGroup", "sdk-Namespace-9080",
                new CheckNameAvailability().withName("sdk-DisasterRecovery-9474"), com.azure.core.util.Context.NONE);
    }
}
