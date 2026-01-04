
import com.azure.resourcemanager.appservice.models.KeyType;
import com.azure.resourcemanager.appservice.models.RegenerateActionParameter;

/**
 * Samples for Workflows RegenerateAccessKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Workflows_RegenerateAccessKey.json
     */
    /**
     * Sample code: Regenerate the callback URL access key for request triggers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        regenerateTheCallbackURLAccessKeyForRequestTriggers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflows().regenerateAccessKeyWithResponse("testResourceGroup",
            "test-name", "testWorkflowName", new RegenerateActionParameter().withKeyType(KeyType.PRIMARY),
            com.azure.core.util.Context.NONE);
    }
}
