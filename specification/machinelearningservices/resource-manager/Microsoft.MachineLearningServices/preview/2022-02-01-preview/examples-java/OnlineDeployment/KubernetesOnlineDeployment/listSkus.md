```java
import com.azure.core.util.Context;

/** Samples for OnlineDeployments ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/KubernetesOnlineDeployment/listSkus.json
     */
    /**
     * Sample code: List Kubernetes Online Deployment Skus.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listKubernetesOnlineDeploymentSkus(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineDeployments()
            .listSkus("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", 1, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.
