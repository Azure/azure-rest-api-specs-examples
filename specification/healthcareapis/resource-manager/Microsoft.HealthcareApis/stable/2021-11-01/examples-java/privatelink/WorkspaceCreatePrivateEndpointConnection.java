import com.azure.core.util.Context;
import com.azure.resourcemanager.healthcareapis.fluent.models.PrivateEndpointConnectionDescriptionInner;
import com.azure.resourcemanager.healthcareapis.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.healthcareapis.models.PrivateLinkServiceConnectionState;

/** Samples for WorkspacePrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceCreatePrivateEndpointConnection.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnection_CreateOrUpdate.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateEndpointConnectionCreateOrUpdate(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .workspacePrivateEndpointConnections()
            .createOrUpdate(
                "testRG",
                "workspace1",
                "myConnection",
                new PrivateEndpointConnectionDescriptionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("Auto-Approved")),
                Context.NONE);
    }
}
