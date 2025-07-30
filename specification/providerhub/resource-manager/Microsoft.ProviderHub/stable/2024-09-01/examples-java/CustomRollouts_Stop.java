
/**
 * Samples for CustomRollouts Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/CustomRollouts_Stop.
     * json
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
