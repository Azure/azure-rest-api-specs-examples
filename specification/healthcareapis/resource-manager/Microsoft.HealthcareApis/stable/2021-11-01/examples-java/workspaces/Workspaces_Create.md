Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.healthcareapis.models.WorkspaceProperties;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_Create.json
     */
    /**
     * Sample code: Create or update a workspace.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void createOrUpdateAWorkspace(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .workspaces()
            .define("workspace1")
            .withExistingResourceGroup("testRG")
            .withRegion("westus")
            .withProperties(new WorkspaceProperties())
            .create();
    }
}
```
