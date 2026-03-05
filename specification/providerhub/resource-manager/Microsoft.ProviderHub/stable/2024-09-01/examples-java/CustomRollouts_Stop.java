
/**
 * Samples for CustomRollouts Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/CustomRollouts_Stop.json
     */
    /**
     * Sample code: CustomRollouts_Stop.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void customRolloutsStop(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.customRollouts().stopWithResponse("Microsoft.Contoso", "2020week10", com.azure.core.util.Context.NONE);
    }
}
