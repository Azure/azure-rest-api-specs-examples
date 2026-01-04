
/**
 * Samples for CredentialSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * CredentialSetGet.json
     */
    /**
     * Sample code: CredentialSetGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void credentialSetGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCredentialSets().getWithResponse("myResourceGroup",
            "myRegistry", "myCredentialSet", com.azure.core.util.Context.NONE);
    }
}
