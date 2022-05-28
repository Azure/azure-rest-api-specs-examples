Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.PublicNetworkAccess;
import com.azure.resourcemanager.machinelearning.models.Workspace;

/** Samples for Workspaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/update.json
     */
    /**
     * Sample code: Update Workspace.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void updateWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        Workspace resource =
            manager
                .workspaces()
                .getByResourceGroupWithResponse("workspace-1234", "testworkspace", Context.NONE)
                .getValue();
        resource
            .update()
            .withDescription("new description")
            .withFriendlyName("New friendly name")
            .withPublicNetworkAccess(PublicNetworkAccess.DISABLED)
            .apply();
    }
}
```
