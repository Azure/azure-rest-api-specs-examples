
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void
        operationsList(com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
