
/**
 * Samples for DeploymentStacks DeleteAtSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackSubscriptionDelete.json
     */
    /**
     * Sample code: DeploymentStacksSubscriptionDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksSubscriptionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .deleteAtSubscription("simpleDeploymentStack", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
