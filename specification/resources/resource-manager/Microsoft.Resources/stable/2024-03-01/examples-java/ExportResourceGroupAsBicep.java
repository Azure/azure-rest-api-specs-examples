
import com.azure.resourcemanager.resources.models.ExportTemplateOutputFormat;
import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/**
 * Samples for ResourceGroups ExportTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * ExportResourceGroupAsBicep.json
     */
    /**
     * Sample code: Export a resource group as Bicep.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportAResourceGroupAsBicep(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getResourceGroups().exportTemplate("my-resource-group",
            new ExportTemplateRequest().withResources(Arrays.asList("*"))
                .withOptions("IncludeParameterDefaultValue,IncludeComments")
                .withOutputFormat(ExportTemplateOutputFormat.BICEP),
            com.azure.core.util.Context.NONE);
    }
}
