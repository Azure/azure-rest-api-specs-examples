/** Samples for DefaultRollouts Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_Stop.json
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
