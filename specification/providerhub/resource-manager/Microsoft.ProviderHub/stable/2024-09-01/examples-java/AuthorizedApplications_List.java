
/**
 * Samples for AuthorizedApplications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/AuthorizedApplications_List.json
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
