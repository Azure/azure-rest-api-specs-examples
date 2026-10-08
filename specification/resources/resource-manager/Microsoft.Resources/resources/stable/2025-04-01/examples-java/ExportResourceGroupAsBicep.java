
import com.azure.resourcemanager.resources.models.ExportTemplateOutputFormat;
import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/**
 * Samples for ResourceGroups ExportTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/ExportResourceGroupAsBicep.json
     */
    /**
     * Sample code: Export a resource group as Bicep.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void exportAResourceGroupAsBicep(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getResourceGroups().exportTemplate("my-resource-group",
            new ExportTemplateRequest().withResources(Arrays.asList("*"))
                .withOptions("IncludeParameterDefaultValue,IncludeComments")
                .withOutputFormat(ExportTemplateOutputFormat.BICEP),
            com.azure.core.util.Context.NONE);
    }
}
