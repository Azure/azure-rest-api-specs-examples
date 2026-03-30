
import com.azure.resourcemanager.appservice.fluent.models.RemotePrivateEndpointConnectionArmResourceInner;
import com.azure.resourcemanager.appservice.models.PrivateLinkConnectionState;

/**
 * Samples for WebApps ApproveOrRejectPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ApproveRejectSitePrivateEndpointConnection.json
     */
    /**
     * Sample code: Approves or rejects a private endpoint connection for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void approvesOrRejectsAPrivateEndpointConnectionForASite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().approveOrRejectPrivateEndpointConnection("rg", "testSite", "connection",
            new RemotePrivateEndpointConnectionArmResourceInner()
                .withPrivateLinkServiceConnectionState(new PrivateLinkConnectionState().withStatus("Approved")
                    .withDescription("Approved by admin.").withActionsRequired("")),
            com.azure.core.util.Context.NONE);
    }
}
