
/**
 * Samples for DefaultRollouts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/DefaultRollouts_Get.json
     */
    /**
     * Sample code: DefaultRollouts_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void defaultRolloutsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.defaultRollouts().getWithResponse("Microsoft.Contoso", "2020week10", com.azure.core.util.Context.NONE);
    }
}
