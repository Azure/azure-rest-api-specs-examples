
import java.util.UUID;

/**
 * Samples for AuthorizedApplications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * AuthorizedApplications_Delete.json
     */
    /**
     * Sample code: AuthorizedApplications_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void authorizedApplicationsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.authorizedApplications().deleteWithResponse("Microsoft.Contoso",
            UUID.fromString("760505bf-dcfa-4311-b890-18da392a00b2"), com.azure.core.util.Context.NONE);
    }
}
