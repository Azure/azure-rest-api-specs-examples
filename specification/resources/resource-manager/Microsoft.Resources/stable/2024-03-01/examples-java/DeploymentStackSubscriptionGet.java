
/**
 * Samples for DeploymentStacks GetAtSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackSubscriptionGet.json
     */
    /**
     * Sample code: DeploymentStacksSubscriptionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksSubscriptionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .getAtSubscriptionWithResponse("simpleDeploymentStack", com.azure.core.util.Context.NONE);
    }
}
