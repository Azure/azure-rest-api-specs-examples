Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkspacePrivateEndpointConnections ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceListPrivateEndpointConnections.json
     */
    /**
     * Sample code: WorkspacePrivateEndpointConnection_List.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateEndpointConnectionList(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspacePrivateEndpointConnections().listByWorkspace("testRG", "workspace1", Context.NONE);
    }
}
```
