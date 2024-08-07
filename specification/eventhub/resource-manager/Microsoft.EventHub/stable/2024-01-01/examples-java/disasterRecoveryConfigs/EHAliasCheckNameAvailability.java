
import com.azure.resourcemanager.eventhubs.models.CheckNameAvailabilityParameter;

/**
 * Samples for DisasterRecoveryConfigs CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/disasterRecoveryConfigs/
     * EHAliasCheckNameAvailability.json
     */
    /**
     * Sample code: NamespacesCheckNameAvailability.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespacesCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getDisasterRecoveryConfigs().checkNameAvailabilityWithResponse(
            "exampleResourceGroup", "sdk-Namespace-9080",
            new CheckNameAvailabilityParameter().withName("sdk-DisasterRecovery-9474"),
            com.azure.core.util.Context.NONE);
    }
}
