Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for WorkspaceConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/create.json
     */
    /**
     * Sample code: CreateWorkspaceConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createWorkspaceConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .workspaceConnections()
            .define("connection-1")
            .withExistingWorkspace("resourceGroup-1", "workspace-1")
            .withCategory("ACR")
            .withTarget("www.facebook.com")
            .withAuthType("PAT")
            .withValue("secrets")
            .create();
    }
}
```
