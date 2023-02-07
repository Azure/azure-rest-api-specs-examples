/** Samples for DefaultRollouts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_Delete.json
     */
    /**
     * Sample code: DefaultRollouts_Delete.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void defaultRolloutsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .defaultRollouts()
            .deleteByResourceGroupWithResponse("Microsoft.Contoso", "2020week10", com.azure.core.util.Context.NONE);
    }
}
