
/**
 * Samples for DefaultRollouts Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/DefaultRollouts_Stop.json
     */
    /**
     * Sample code: DefaultRollouts_Stop.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void defaultRolloutsStop(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.defaultRollouts().stopWithResponse("Microsoft.Contoso", "2020week10", com.azure.core.util.Context.NONE);
    }
}
