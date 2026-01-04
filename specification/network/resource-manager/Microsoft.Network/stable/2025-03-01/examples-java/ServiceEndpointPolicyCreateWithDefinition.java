
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyDefinitionInner;
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyInner;
import java.util.Arrays;

/**
 * Samples for ServiceEndpointPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ServiceEndpointPolicyCreateWithDefinition.json
     */
    /**
     * Sample code: Create service endpoint policy with definition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServiceEndpointPolicyWithDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicies().createOrUpdate("rg1", "testPolicy",
            new ServiceEndpointPolicyInner().withLocation("westus")
                .withServiceEndpointPolicyDefinitions(Arrays.asList(new ServiceEndpointPolicyDefinitionInner()
                    .withName("StorageServiceEndpointPolicyDefinition")
                    .withDescription("Storage Service EndpointPolicy Definition").withService("Microsoft.Storage")
                    .withServiceResources(Arrays.asList("/subscriptions/subid1",
                        "/subscriptions/subid1/resourceGroups/storageRg",
                        "/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")))),
            com.azure.core.util.Context.NONE);
    }
}
