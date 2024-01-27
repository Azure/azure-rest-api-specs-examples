
/**
 * Samples for Apms ListSecretKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Apms_ListSecretKeys.
     * json
     */
    /**
     * Sample code: Apms_ListSecretKeys.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apmsListSecretKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApms().listSecretKeysWithResponse("myResourceGroup",
            "myservice", "myappinsights", com.azure.core.util.Context.NONE);
    }
}
