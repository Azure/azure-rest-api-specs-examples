/** Samples for ProviderOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ProviderOperations_List.json
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
