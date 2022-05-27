Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyDefinitionInner;
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyInner;
import java.util.Arrays;

/** Samples for ServiceEndpointPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ServiceEndpointPolicyCreateWithDefinition.json
     */
    /**
     * Sample code: Create service endpoint policy with definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServiceEndpointPolicyWithDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getServiceEndpointPolicies()
            .createOrUpdate(
                "rg1",
                "testPolicy",
                new ServiceEndpointPolicyInner()
                    .withLocation("westus")
                    .withServiceEndpointPolicyDefinitions(
                        Arrays
                            .asList(
                                new ServiceEndpointPolicyDefinitionInner()
                                    .withName("StorageServiceEndpointPolicyDefinition")
                                    .withDescription("Storage Service EndpointPolicy Definition")
                                    .withService("Microsoft.Storage")
                                    .withServiceResources(
                                        Arrays
                                            .asList(
                                                "/subscriptions/subid1",
                                                "/subscriptions/subid1/resourceGroups/storageRg",
                                                "/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")))),
                Context.NONE);
    }
}
```
