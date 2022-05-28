Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for OnlineDeployments ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/ManagedOnlineDeployment/listSkus.json
     */
    /**
     * Sample code: List Managed Online Deployment Skus.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listManagedOnlineDeploymentSkus(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .onlineDeployments()
            .listSkus("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", 1, null, Context.NONE);
    }
}
```
