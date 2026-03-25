
/**
 * Samples for CredentialSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/CredentialSetGet.json
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
