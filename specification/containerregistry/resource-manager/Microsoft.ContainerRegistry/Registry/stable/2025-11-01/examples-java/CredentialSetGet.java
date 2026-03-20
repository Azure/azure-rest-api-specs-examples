
/**
 * Samples for CredentialSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CredentialSetGet.json
     */
    /**
     * Sample code: CredentialSetGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void credentialSetGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCredentialSets().getWithResponse("myResourceGroup", "myRegistry", "myCredentialSet",
            com.azure.core.util.Context.NONE);
    }
}
