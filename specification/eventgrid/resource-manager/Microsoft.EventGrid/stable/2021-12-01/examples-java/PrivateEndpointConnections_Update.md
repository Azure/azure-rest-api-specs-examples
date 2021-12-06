Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.4/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.eventgrid.models.ConnectionState;
import com.azure.resourcemanager.eventgrid.models.PersistedConnectionStatus;
import com.azure.resourcemanager.eventgrid.models.PrivateEndpointConnectionsParentType;

/** Samples for PrivateEndpointConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/PrivateEndpointConnections_Update.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void privateEndpointConnectionsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .privateEndpointConnections()
            .update(
                "examplerg",
                PrivateEndpointConnectionsParentType.TOPICS,
                "exampletopic1",
                "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new ConnectionState()
                            .withStatus(PersistedConnectionStatus.APPROVED)
                            .withDescription("approving connection")
                            .withActionsRequired("None")),
                Context.NONE);
    }
}
```
