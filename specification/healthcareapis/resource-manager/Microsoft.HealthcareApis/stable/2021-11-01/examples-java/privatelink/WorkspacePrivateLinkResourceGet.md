Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkspacePrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspacePrivateLinkResourceGet.json
     */
    /**
     * Sample code: WorkspacePrivateLinkResources_Get.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void workspacePrivateLinkResourcesGet(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .workspacePrivateLinkResources()
            .getWithResponse("testRG", "workspace1", "healthcareworkspace", Context.NONE);
    }
}
```
