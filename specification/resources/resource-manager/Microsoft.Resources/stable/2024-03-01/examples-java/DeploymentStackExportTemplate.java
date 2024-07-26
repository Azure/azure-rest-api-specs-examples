
/**
 * Samples for DeploymentStacks ExportTemplateAtResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackExportTemplate.json
     */
    /**
     * Sample code: DeploymentStacksResourceGroupExportTemplate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deploymentStacksResourceGroupExportTemplate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .exportTemplateAtResourceGroupWithResponse("deploymentStacksRG", "simpleDeploymentStack",
                com.azure.core.util.Context.NONE);
    }
}
