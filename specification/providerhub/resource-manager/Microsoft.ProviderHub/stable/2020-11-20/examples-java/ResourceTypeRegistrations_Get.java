/** Samples for ResourceTypeRegistrations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_Get.json
     */
    /**
     * Sample code: ResourceTypeRegistrations_Get.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void resourceTypeRegistrationsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .resourceTypeRegistrations()
            .getWithResponse("Microsoft.Contoso", "employees", com.azure.core.util.Context.NONE);
    }
}
