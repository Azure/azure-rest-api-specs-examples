
import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/**
 * Samples for ResourceGroups ExportTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/ExportResourceGroup.json
     */
    /**
     * Sample code: Export a resource group.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void exportAResourceGroup(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getResourceGroups().exportTemplate("my-resource-group", new ExportTemplateRequest()
            .withResources(Arrays.asList("*")).withOptions("IncludeParameterDefaultValue,IncludeComments"),
            com.azure.core.util.Context.NONE);
    }
}
