
import java.util.UUID;

/**
 * Samples for AuthorizedApplications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * AuthorizedApplications_Get.json
     */
    /**
     * Sample code: AuthorizedApplications_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void authorizedApplicationsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.authorizedApplications().getWithResponse("Microsoft.Contoso",
            UUID.fromString("760505bf-dcfa-4311-b890-18da392a00b2"), com.azure.core.util.Context.NONE);
    }
}
