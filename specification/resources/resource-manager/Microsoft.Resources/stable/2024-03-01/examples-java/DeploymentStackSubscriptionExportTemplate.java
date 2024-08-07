
/**
 * Samples for DeploymentStacks ExportTemplateAtSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackSubscriptionExportTemplate.json
     */
    /**
     * Sample code: DeploymentStacksSubscriptionExportTemplate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deploymentStacksSubscriptionExportTemplate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .exportTemplateAtSubscriptionWithResponse("simpleDeploymentStack", com.azure.core.util.Context.NONE);
    }
}
