
import com.azure.core.util.Context;

/** Samples for DisasterRecoveryConfigs List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/
     * EHAliasList.json
     */
    /**
     * Sample code: EHAliasList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eHAliasList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getDisasterRecoveryConfigs().list("exampleResourceGroup",
            "sdk-Namespace-8859", Context.NONE);
    }
}
