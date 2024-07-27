
/**
 * Samples for DeploymentStacks ExportTemplateAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * DeploymentStackManagementGroupExportTemplate.json
     */
    /**
     * Sample code: DeploymentStacksManagementGroupExportTemplate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deploymentStacksManagementGroupExportTemplate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().deploymentStackClient().getDeploymentStacks()
            .exportTemplateAtManagementGroupWithResponse("myMg", "simpleDeploymentStack",
                com.azure.core.util.Context.NONE);
    }
}
