
import com.azure.resourcemanager.appservice.fluent.models.RemotePrivateEndpointConnectionArmResourceInner;
import com.azure.resourcemanager.appservice.models.PrivateLinkConnectionState;

/**
 * Samples for StaticSites ApproveOrRejectPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ApproveRejectSitePrivateEndpointConnection.json
     */
    /**
     * Sample code: Approves or rejects a private endpoint connection for a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        approvesOrRejectsAPrivateEndpointConnectionForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().approveOrRejectPrivateEndpointConnection("rg",
            "testSite", "connection",
            new RemotePrivateEndpointConnectionArmResourceInner()
                .withPrivateLinkServiceConnectionState(new PrivateLinkConnectionState().withStatus("Approved")
                    .withDescription("Approved by admin.").withActionsRequired("")),
            com.azure.core.util.Context.NONE);
    }
}
