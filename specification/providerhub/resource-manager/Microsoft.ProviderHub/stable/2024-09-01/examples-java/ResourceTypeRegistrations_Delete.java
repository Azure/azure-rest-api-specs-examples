
/**
 * Samples for ResourceTypeRegistrations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ResourceTypeRegistrations_Delete.json
     */
    /**
     * Sample code: ResourceTypeRegistrations_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        resourceTypeRegistrationsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.resourceTypeRegistrations().delete("Microsoft.Contoso", "testResourceType",
            com.azure.core.util.Context.NONE);
    }
}
