
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EHOperations_List.json
     */
    /**
     * Sample code: EHOperations_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eHOperationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
