
/**
 * Samples for AuthorizedApplications List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * AuthorizedApplications_List.json
     */
    /**
     * Sample code: AuthorizedApplications_List.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void authorizedApplicationsList(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.authorizedApplications().list("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
