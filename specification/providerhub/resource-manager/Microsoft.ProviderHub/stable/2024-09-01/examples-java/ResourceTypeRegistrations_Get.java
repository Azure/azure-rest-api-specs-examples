
/**
 * Samples for ResourceTypeRegistrations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ResourceTypeRegistrations_Get.json
     */
    /**
     * Sample code: ResourceTypeRegistrations_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void resourceTypeRegistrationsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.resourceTypeRegistrations().getWithResponse("Microsoft.Contoso", "employees",
            com.azure.core.util.Context.NONE);
    }
}
