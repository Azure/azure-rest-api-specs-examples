
/**
 * Samples for ProviderOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ProviderOperations_List.json
     */
    /**
     * Sample code: ProviderOperations_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void providerOperationsList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.providerOperations().list(com.azure.core.util.Context.NONE);
    }
}
