import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/examples/MsiOperationsList.json
     */
    /**
     * Sample code: MsiOperationsList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void msiOperationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getOperations().list(Context.NONE);
    }
}
