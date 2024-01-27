
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/SBOperations_List.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void operationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
