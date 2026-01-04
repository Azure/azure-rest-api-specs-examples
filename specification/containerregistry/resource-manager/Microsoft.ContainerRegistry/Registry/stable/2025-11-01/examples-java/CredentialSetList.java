
/**
 * Samples for CredentialSets List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * CredentialSetList.json
     */
    /**
     * Sample code: CredentialSetList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void credentialSetList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCredentialSets().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
