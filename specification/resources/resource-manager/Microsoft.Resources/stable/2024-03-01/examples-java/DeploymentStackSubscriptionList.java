
/**
 * Samples for DeploymentStacks List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackSubscriptionList.json
     */
    /**
     * Sample code: DeploymentStacksSubscriptionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentStacksSubscriptionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .list(com.azure.core.util.Context.NONE);
    }
}
