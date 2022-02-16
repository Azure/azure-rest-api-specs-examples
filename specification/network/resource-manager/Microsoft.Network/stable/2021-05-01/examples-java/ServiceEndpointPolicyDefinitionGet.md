Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ServiceEndpointPolicyDefinitions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ServiceEndpointPolicyDefinitionGet.json
     */
    /**
     * Sample code: Get service endpoint definition in service endpoint policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getServiceEndpointDefinitionInServiceEndpointPolicy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getServiceEndpointPolicyDefinitions()
            .getWithResponse("rg1", "testPolicy", "testDefinition", Context.NONE);
    }
}
```
