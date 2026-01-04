
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyDefinitionInner;
import java.util.Arrays;

/**
 * Samples for ServiceEndpointPolicyDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ServiceEndpointPolicyDefinitionCreate.json
     */
    /**
     * Sample code: Create service endpoint policy definition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServiceEndpointPolicyDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceEndpointPolicyDefinitions()
            .createOrUpdate("rg1", "testPolicy", "testDefinition", new ServiceEndpointPolicyDefinitionInner()
                .withDescription("Storage Service EndpointPolicy Definition").withService("Microsoft.Storage")
                .withServiceResources(Arrays.asList("/subscriptions/subid1",
                    "/subscriptions/subid1/resourceGroups/storageRg",
                    "/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")),
                com.azure.core.util.Context.NONE);
    }
}
