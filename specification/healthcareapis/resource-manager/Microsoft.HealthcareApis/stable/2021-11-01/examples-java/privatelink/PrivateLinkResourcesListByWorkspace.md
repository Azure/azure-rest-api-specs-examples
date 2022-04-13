Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkspacePrivateLinkResources ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/PrivateLinkResourcesListByWorkspace.json
     */
    /**
     * Sample code: WorkspacePrivateLinkResources_ListGroupIds.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateLinkResourcesListGroupIds(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.workspacePrivateLinkResources().listByWorkspace("testRG", "workspace1", Context.NONE);
    }
}
```
