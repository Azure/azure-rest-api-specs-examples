import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void operationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
