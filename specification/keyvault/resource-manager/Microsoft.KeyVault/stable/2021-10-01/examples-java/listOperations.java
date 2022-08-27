import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/listOperations.json
     */
    /**
     * Sample code: Lists available Rest API operations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAvailableRestAPIOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
