
/**
 * Samples for CredentialSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/CredentialSetList.json
     */
    /**
     * Sample code: CredentialSetList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void credentialSetList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCredentialSets().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
