import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/OperationList.json
     */
    /**
     * Sample code: Get a list of operations for a resource provider.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAListOfOperationsForAResourceProvider(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
