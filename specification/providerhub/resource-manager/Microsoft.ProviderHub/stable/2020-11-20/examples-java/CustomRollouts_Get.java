/** Samples for CustomRollouts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/CustomRollouts_Get.json
     */
    /**
     * Sample code: CustomRollouts_Get.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void customRolloutsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .customRollouts()
            .getWithResponse("Microsoft.Contoso", "canaryTesting99", com.azure.core.util.Context.NONE);
    }
}
