
/**
 * Samples for ManagedHsmKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/managedHsmListKeys.json
     */
    /**
     * Sample code: List keys in the managed HSM.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listKeysInTheManagedHSM(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsmKeys().list("sample-group", "sample-managedhsm-name",
            com.azure.core.util.Context.NONE);
    }
}
